package challenge

import (
	"context"
	"time"

	"github.com/0chain/blobber/code/go/0chain.net/blobbercore/models"
	"github.com/0chain/errors"
	"github.com/0chain/gosdk/constants"
	"gorm.io/gorm"
)

// Exists check challenge if it is added
func Exists(ctx context.Context, db *gorm.DB, challengeID string) (bool, error) {

	if len(challengeID) == 0 {
		return false, errors.Throw(constants.ErrInvalidParameter, "challengeID")
	}

	var count int64
	err := db.Table(models.TableNameChallenge).Where(SQLWhereId, challengeID).Count(&count).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}

		return false, errors.ThrowLog(err.Error(), constants.ErrBadDatabaseOperation)
	}

	return count > 0, nil
}

// getTodoChallenges get todo challenges list which status is accepted
func getTodoChallenges(ctx context.Context, db *gorm.DB) ([]*models.Challenge, error) {

	var list []*models.Challenge
	err := db.Table(models.TableNameChallenge).
		Where(SQLWhereTodo).
		Order("sequence desc").
		Find(&list).Error

	if err != nil {
		return nil, errors.ThrowLog(err.Error(), constants.ErrBadDatabaseOperation)
	}

	return list, nil
}

// cancelChallenge cancel invalid challenge from blockchain
func cancelChallenge(ctx context.Context, db *gorm.DB, id, message string) error {

	return db.Table(models.TableNameChallenge).
		Where(SQLWhereId, id).
		Select("status", "status_messsage", "updated_at").
		Updates(map[string]interface{}{
			"status":          models.ChallengeStatusSkipped,
			"status_messsage": message,
			"updated_at":      time.Now().Unix(),
		}).Error

}

const (
	SQLWhereId   = `challenges.challenge_id = ?`
	SQLWhereTodo = `challenges.status = 1`
)

// DryRun  Creates a prepared statement when executing any SQL and caches them to speed up future calls
// https://gorm.io/docs/performance.html#Caches-Prepared-Statement
func DryRun(db *gorm.DB) {

	// https://gorm.io/docs/session.html#DryRun
	// Session mode
	tx := db.Session(&gorm.Session{PrepareStmt: true, DryRun: true})

	// use Table instead of Model to reduce reflect times

	Exists(context.TODO(), tx, "id")
	getTodoChallenges(context.TODO(), tx)
}
