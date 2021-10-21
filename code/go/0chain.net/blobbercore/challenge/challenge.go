package challenge

import (
	"context"
	"encoding/json"
	"time"

	"github.com/0chain/blobber/code/go/0chain.net/blobbercore/config"
	"github.com/0chain/blobber/code/go/0chain.net/blobbercore/datastore"
	"github.com/0chain/blobber/code/go/0chain.net/blobbercore/models"
	"github.com/0chain/blobber/code/go/0chain.net/core/chain"
	"github.com/0chain/blobber/code/go/0chain.net/core/common"
	"github.com/0chain/blobber/code/go/0chain.net/core/lock"
	"github.com/0chain/blobber/code/go/0chain.net/core/node"
	"github.com/0chain/blobber/code/go/0chain.net/core/transaction"
	"github.com/remeh/sizedwaitgroup"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/0chain/blobber/code/go/0chain.net/core/logging"
)

type BCChallengeResponse struct {
	BlobberID  string             `json:"blobber_id"`
	Challenges []*ChallengeEntity `json:"challenges"`
}

func getChallenges(ctx context.Context) *BlobberChallenge {
	defer func() {
		if r := recover(); r != nil {
			logging.Logger.Error("[recover]challenge", zap.Any("err", r))
		}
	}()

	params := make(map[string]string)
	params["blobber"] = node.Self.ID

	blobberChallenges := &BlobberChallenge{}
	blobberChallenges.Challenges = make([]*StorageChallenge, 0)
	buf, err := transaction.MakeSCRestAPICall(transaction.STORAGE_CONTRACT_ADDRESS, "/openchallenges", params, chain.GetServerChain())
	if err != nil {
		logging.Logger.Error("[challenge]get: ", zap.Error(err))
		return nil
	}

	err = json.Unmarshal(buf, blobberChallenges)
	if err != nil {
		logging.Logger.Error("[challenge]json: ", zap.Error(err))
		return nil
	}

	return blobberChallenges

}

// acceptChallenges get challenge from blockchain , and add them in database if it doesn't exists
func acceptChallenges(ctx context.Context) {
	defer func() {
		if r := recover(); r != nil {
			logging.Logger.Error("[recover]challenge", zap.Any("err", r))
		}
	}()

	blobberChallenges := getChallenges(ctx)

	if blobberChallenges == nil {
		logging.Logger.Info("[challenge]No open challenge")
		return
	}

	db := datastore.GetStore().GetDB()

	for _, c := range blobberChallenges.Challenges {
		if c == nil || len(c.ID) == 0 {
			continue
		}

		ok, err := Exists(ctx, db, c.ID)
		if err != nil {
			logging.Logger.Error("[challenge]exists: ", zap.Error(err))
			return
		}

		if ok {
			continue
		}

		logging.Logger.Info("[challenge]accept " + c.ID)

		it := &models.Challenge{
			ChallengeID: c.ID,

			AllocationID:   c.AllocationID,
			AllocationRoot: c.AllocationRoot,

			Status: models.ChallengeStatusAccepted,
			Result: models.ChallengeResultUnknown,

			Seed:       c.RandomNumber,
			Validators: models.ToJSON(c.Validators),

			CreatedAt: common.ToTime(c.Created),
			UpdatedAt: time.Now().UTC(),
		}

		err = db.Transaction(func(tx *gorm.DB) error {
			return db.Create(it).Error
		})

		if err != nil {
			logging.Logger.Error("[challenge]tx: ", zap.Error(err))
			return
		}
	}

}

// processAccepted read accepted challenge from db, and send them to validator to pass challenge
func processAccepted(ctx context.Context) {
	defer func() {
		if r := recover(); r != nil {
			logging.Logger.Error("[recover]challenge", zap.Any("err", r))
		}
	}()
	rctx := datastore.GetStore().CreateTransaction(ctx)
	db := datastore.GetStore().GetTransaction(rctx)
	openchallenges := make([]*ChallengeEntity, 0)
	db.Where(ChallengeEntity{Status: Accepted}).Find(&openchallenges)
	if len(openchallenges) > 0 {
		swg := sizedwaitgroup.New(config.Configuration.ChallengeResolveNumWorkers)
		for _, openchallenge := range openchallenges {
			logging.Logger.Info("Processing the challenge", zap.Any("challenge_id", openchallenge.ChallengeID), zap.Any("openchallenge", openchallenge))
			err := openchallenge.UnmarshalFields()
			if err != nil {
				logging.Logger.Error("Error unmarshaling challenge entity.", zap.Error(err))
				continue
			}
			swg.Add()
			go func(redeemCtx context.Context, challengeEntity *ChallengeEntity) {
				redeemCtx = datastore.GetStore().CreateTransaction(redeemCtx)
				defer redeemCtx.Done()
				err := loadValidationTickets(redeemCtx, challengeEntity)
				if err != nil {
					logging.Logger.Error("Getting validation tickets failed", zap.Any("challenge_id", challengeEntity.ChallengeID), zap.Error(err))
				}
				db := datastore.GetStore().GetTransaction(redeemCtx)
				err = db.Commit().Error
				if err != nil {
					logging.Logger.Error("Error commiting the readmarker redeem", zap.Error(err))
				}
				swg.Done()
			}(ctx, openchallenge)
		}
		swg.Wait()
	}
	db.Rollback()
	rctx.Done()
}

// loadValidationTickets load validation tickets for challenge
func loadValidationTickets(ctx context.Context, challengeObj *ChallengeEntity) error {
	mutex := lock.GetMutex(challengeObj.TableName(), challengeObj.ChallengeID)
	mutex.Lock()

	defer func() {
		if r := recover(); r != nil {
			logging.Logger.Error("[recover] LoadValidationTickets", zap.Any("err", r))
		}
	}()

	err := challengeObj.LoadValidationTickets(ctx)
	if err != nil {
		logging.Logger.Error("Error getting the validation tickets", zap.Error(err), zap.String("challenge_id", challengeObj.ChallengeID))
	}

	return err
}

func commitProcessed(ctx context.Context) {
	defer func() {
		if r := recover(); r != nil {
			logging.Logger.Error("[recover]challenge", zap.Any("err", r))
		}
	}()

	rctx := datastore.GetStore().CreateTransaction(ctx)
	db := datastore.GetStore().GetTransaction(rctx)
	openchallenges := make([]*ChallengeEntity, 0)

	db.Where(ChallengeEntity{Status: Processed}).
		Order("sequence").
		Find(&openchallenges)

	for _, openchallenge := range openchallenges {
		logging.Logger.Info("Attempting to commit challenge", zap.Any("challenge_id", openchallenge.ChallengeID), zap.Any("openchallenge", openchallenge))
		if err := openchallenge.UnmarshalFields(); err != nil {
			logging.Logger.Error("ChallengeEntity_UnmarshalFields", zap.String("challenge_id", openchallenge.ChallengeID), zap.Error(err))
		}
		mutex := lock.GetMutex(openchallenge.TableName(), openchallenge.ChallengeID)
		mutex.Lock()
		redeemCtx := datastore.GetStore().CreateTransaction(ctx)
		err := openchallenge.CommitChallenge(redeemCtx, false)
		if err != nil {
			logging.Logger.Error("Error committing to blockchain",
				zap.Error(err),
				zap.String("challenge_id", openchallenge.ChallengeID))
		}
		mutex.Unlock()
		db := datastore.GetStore().GetTransaction(redeemCtx)
		db.Commit()
		if err == nil && openchallenge.Status == Committed {
			logging.Logger.Info("Challenge has been submitted to blockchain",
				zap.Any("id", openchallenge.ChallengeID),
				zap.String("txn", openchallenge.CommitTxnID))
		} else {
			logging.Logger.Info("Challenge was not committed", zap.Any("challenge_id", openchallenge.ChallengeID))
			break
		}
	}

	db.Rollback()
	rctx.Done()
}
