package challenge

import (
	"context"

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
	err := db.Table(models.TableNameChallenge).Where(SQLWhereExists, challengeID).Count(&count).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}

		return false, errors.ThrowLog(err.Error(), constants.ErrBadDatabaseOperation)
	}

	return count > 0, nil
}

// GetLatestChallengeID get latest challenge_id order by sequence
func GetLatestChallengeID(ctx context.Context, db *gorm.DB) (string, error) {

	var result map[string]string
	err := db.Table(models.TableNameChallenge).Select("challenge_id").Order("sequence desc").First(&result).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", nil
		}

		return "", errors.ThrowLog(err.Error(), constants.ErrBadDatabaseOperation)
	}

	return result["challenge_id"], nil
}

const (
	SQLWhereExists = `challenges.challenge_id = ?`
)

// DryRun  Creates a prepared statement when executing any SQL and caches them to speed up future calls
// https://gorm.io/docs/performance.html#Caches-Prepared-Statement
func DryRun(db *gorm.DB) {

	// https://gorm.io/docs/session.html#DryRun
	// Session mode
	tx := db.Session(&gorm.Session{PrepareStmt: true, DryRun: true})

	// use Table instead of Model to reduce reflect times

	// prepare statement for Exists
	var count int64
	tx.Table(models.TableNameChallenge).Where(SQLWhereExists, "id").Count(&count)

	// prepare statement for GetLatestChallengeID
	var result map[string]string
	db.Table(models.TableNameChallenge).Select("challenge_id").Order("sequence desc").First(&result)

}
