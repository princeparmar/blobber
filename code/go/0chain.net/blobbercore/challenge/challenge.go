package challenge

import (
	"context"
	"encoding/json"
	"errors"
	"math"
	"math/rand"
	"strconv"
	"time"

	"github.com/0chain/blobber/code/go/0chain.net/blobbercore/allocation"
	"github.com/0chain/blobber/code/go/0chain.net/blobbercore/datastore"
	"github.com/0chain/blobber/code/go/0chain.net/blobbercore/filestore"
	"github.com/0chain/blobber/code/go/0chain.net/blobbercore/models"
	"github.com/0chain/blobber/code/go/0chain.net/blobbercore/reference"
	"github.com/0chain/blobber/code/go/0chain.net/blobbercore/writemarker"
	"github.com/0chain/blobber/code/go/0chain.net/core/chain"
	"github.com/0chain/blobber/code/go/0chain.net/core/common"
	"github.com/0chain/blobber/code/go/0chain.net/core/lock"
	"github.com/0chain/blobber/code/go/0chain.net/core/node"
	"github.com/0chain/blobber/code/go/0chain.net/core/transaction"
	"github.com/0chain/gosdk/constants"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/0chain/blobber/code/go/0chain.net/core/logging"
)

type BCChallengeResponse struct {
	BlobberID  string             `json:"blobber_id"`
	Challenges []*ChallengeEntity `json:"challenges"`
}

func getChallenges(ctx context.Context) *BlobberChallenge {

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

	db := datastore.GetStore().GetDB()

	list, err := getTodoChallenges(ctx, db)
	if err != nil {
		logging.Logger.Error("[challenge]exists: ", zap.Error(err))
		return
	}

	for _, it := range list {

		var validators []ValidationNode

		err := models.FromJSON(&it.Validators, &validators)

		if err != nil {
			cancelChallenge(ctx, db, it.ChallengeID, err.Error())
			continue
		}

		if len(validators) == 0 {
			cancelChallenge(ctx, db, it.ChallengeID, "No validators assigned to the challenge")
			continue
		}

		challengeTask := &ChallengeTask{
			ChallengeID: it.ChallengeID,
		}

		challengeTask.SelectedFile, challengeTask.SelectedBlockIndex, err = getChallengeFile(ctx, db, it.Seed, it.AllocationID)
		if err != nil {
			// cancel it if it is an unknown error
			if !errors.Is(err, constants.ErrBadDatabaseOperation) {
				cancelChallenge(ctx, db, it.ChallengeID, "invalid allocation id "+it.AllocationID)
			}
			//skip it if it is an unknown db error
			continue
		}

		allocationRoot, err := allocation.GetAllocationRoot(ctx, db, it.AllocationID)
		if err != nil {
			// cancel it if it is an known error
			if !errors.Is(err, constants.ErrBadDatabaseOperation) {
				cancelChallenge(ctx, db, it.ChallengeID, "invalid allocation id "+it.AllocationID)
			}
			//skip it if it is an unknown db error
			continue
		}

		challengeTask.WriteMarkers, err = writemarker.GetChallengeWriteMarker(ctx, db, it.AllocationID, it.AllocationRoot, allocationRoot)
		if err != nil {
			// cancel it if it is an unknown error
			if !errors.Is(err, constants.ErrBadDatabaseOperation) {
				cancelChallenge(ctx, db, it.ChallengeID, "invalid allocation id "+it.AllocationID)
			}
			//skip it if it is an unknown db error
			continue
		}

		inputData := &filestore.FileInputData{}
		inputData.Name = challengeTask.SelectedFile.Name
		inputData.Path = challengeTask.SelectedFile.Path
		inputData.Hash = challengeTask.SelectedFile.ContentHash
		inputData.ChunkSize = challengeTask.SelectedFile.ChunkSize

		maxNumBlocks := 1024

		// the file is too small, some of 1024 blocks is not filled
		if challengeTask.SelectedFile.Size < challengeTask.SelectedFile.ChunkSize {
			merkleChunkSize := challengeTask.SelectedFile.ChunkSize / 1024
			maxNumBlocks = int(math.Ceil(float64(challengeTask.SelectedFile.Size) / float64(merkleChunkSize)))
		}

		r := rand.New(rand.NewSource(it.Seed))
		blockIndex := r.Intn(maxNumBlocks)
		selectedBlockBytes, mt, err := filestore.GetFileStore().GetFileBlockForChallenge(it.AllocationID, inputData, blockIndex)

		if err != nil {
			cancelChallenge(ctx, db, it.ChallengeID, "invalid allocation id "+it.AllocationID)
			continue
		}

		challengeTask.SelectedBlockBytes = selectedBlockBytes
		challengeTask.SelectedMerklePath = mt.GetPathByIndex(blockIndex)

		// 	postDataBytes, err := json.Marshal(postData)
		// if err != nil {
		// 	Logger.Error("Error in marshalling the post data for validation. " + err.Error())
		// 	cr.ErrorChallenge(ctx, err)
		// 	return err
		// }
		// responses := make(map[string]ValidationTicket)
		// if cr.ValidationTickets == nil {
		// 	cr.ValidationTickets = make([]*ValidationTicket, len(cr.Validators))
		// }
		// for i, validator := range cr.Validators {
		// 	if cr.ValidationTickets[i] != nil {
		// 		exisitingVT := cr.ValidationTickets[i]
		// 		if len(exisitingVT.Signature) > 0 && exisitingVT.ChallengeID == cr.ChallengeID {
		// 			continue
		// 		}
		// 	}

		// 	url := validator.URL + VALIDATOR_URL

		// 	resp, err := util.SendPostRequest(url, postDataBytes, nil)
		// 	if err != nil {
		// 		Logger.Info("Got error from the validator.", zap.Any("error", err.Error()))
		// 		delete(responses, validator.ID)
		// 		cr.ValidationTickets[i] = nil
		// 		continue
		// 	}
		// 	var validationTicket ValidationTicket
		// 	err = json.Unmarshal(resp, &validationTicket)
		// 	if err != nil {
		// 		Logger.Info("Got error decoding from the validator response .", zap.Any("resp", string(resp)), zap.Any("error", err.Error()))
		// 		delete(responses, validator.ID)
		// 		cr.ValidationTickets[i] = nil
		// 		continue
		// 	}
		// 	Logger.Info("Got response from the validator.", zap.Any("validator_response", validationTicket))
		// 	verified, err := validationTicket.VerifySign()
		// 	if err != nil || !verified {
		// 		Logger.Info("Validation ticket from validator could not be verified.")
		// 		delete(responses, validator.ID)
		// 		cr.ValidationTickets[i] = nil
		// 		continue
		// 	}
		// 	responses[validator.ID] = validationTicket
		// 	cr.ValidationTickets[i] = &validationTicket
		// }

		// numSuccess := 0
		// numFailure := 0

		// numValidatorsResponded := 0
		// for _, vt := range cr.ValidationTickets {
		// 	if vt != nil {
		// 		if vt.Result {
		// 			numSuccess++
		// 		} else {
		// 			numFailure++
		// 		}
		// 		numValidatorsResponded++
		// 	}
		// }

		// Logger.Info("validator response stats", zap.Any("challenge_id", cr.ChallengeID), zap.Any("validator_responses", responses))
		// if numSuccess > (len(cr.Validators)/2) || numFailure > (len(cr.Validators)/2) || numValidatorsResponded == len(cr.Validators) {
		// 	if numSuccess > (len(cr.Validators) / 2) {
		// 		cr.Result = ChallengeSuccess
		// 	} else {
		// 		cr.Result = ChallengeFailure
		// 		//Logger.Error("Challenge failed by the validators", zap.Any("block_num", cr.BlockNum), zap.Any("object_path", objectPath), zap.Any("challenge", cr))
		// 	}

		// 	cr.Status = Processed
		// } else {
		// 	cr.ErrorChallenge(ctx, common.NewError("no_consensus_challenge", "No Consensus on the challenge result. Erroring out the challenge"))
		// 	return common.NewError("no_consensus_challenge", "No Consensus on the challenge result. Erroring out the challenge")
		// }

		// return cr.Save(ctx)

	}

}

// getChallengeFile pick storage file for challenge based on randSeed
func getChallengeFile(ctx context.Context, db *gorm.DB, randSeed int64, allocationID string) (*models.ChallengeStorage, int64, error) {
	root, err := reference.GetChallengeMeta(ctx, db, allocationID)

	if err != nil {
		return nil, 0, err
	}

	if root.NumBlocks < 1 {
		return nil, 0, errors.New("blank allocation")

	}

	obj := &ChallengeTask{
		RootHash: root.Hash,
		Storages: make([][]*models.ChallengeStorage, 0),
	}

	r := rand.New(rand.NewSource(randSeed))
	blockNum := r.Int63n(root.NumBlocks) + 1 //index from 1 instead of 0
	remainingBlocks := blockNum

	selectedPath := "/"

	for level := 0; ; level++ {
		blocks, err := reference.GetNextBlocks(ctx, db, allocationID, selectedPath)

		if err != nil {
			return nil, 0, err
		}

		if len(blocks) == 0 {
			return nil, 0, errors.New("invalid blockNum: " + strconv.FormatInt(blockNum, 10) + "/" + strconv.FormatInt(root.NumBlocks, 10))
		}

		list := make([]*models.ChallengeStorage, 0, len(blocks))
		obj.Storages = append(obj.Storages, list)

		for _, it := range blocks {

			list = append(list, it)

			// it is dir, we need to load its files
			if it.Type == reference.DIRECTORY {
				// block will be this dir's files
				if it.NumBlocks > remainingBlocks {
					selectedPath = it.Path
					continue
				}
			}

			if it.NumBlocks > remainingBlocks {
				// block is found on this file for challenge
				return it, remainingBlocks, nil

			}

			// block is still not found, move to next file/folder
			remainingBlocks -= it.NumBlocks
		}
	}

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
