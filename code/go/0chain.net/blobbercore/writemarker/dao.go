package writemarker

import (
	"context"

	"github.com/0chain/blobber/code/go/0chain.net/blobbercore/models"
	"github.com/0chain/errors"
	"github.com/0chain/gosdk/constants"
	"gorm.io/gorm"
)

// GetChallengeWriteMarker get peanding writemarkers for challenge
func GetChallengeWriteMarker(ctx context.Context, db *gorm.DB, allocationID string, startAllocationRoot string, endAllocationRoot string) ([]*models.ChallengeWriteMarker, error) {

	var seqRange []int64
	err := db.Table(models.TableNameWriteMarker).
		Select("sequence").
		Where(SQLWhereAllocationIdAndRootRange, allocationID, startAllocationRoot, endAllocationRoot).
		Order("sequence").
		Find(&seqRange).Error

	if err != nil {
		return nil, errors.ThrowLog(err.Error(), constants.ErrBadDatabaseOperation)
	}

	if len(seqRange) == 0 {
		return nil, errors.Throw(constants.ErrEntityNotFound, "write_marker_not_found in range allocation_id: "+allocationID, " start_root: "+startAllocationRoot+" end_root: "+endAllocationRoot)
	}

	startSequence := seqRange[0]
	endSequence := startSequence

	if len(seqRange) > 1 {
		endSequence = seqRange[1]
	}

	list := make([]*models.ChallengeWriteMarker, 0)

	err = db.Table(models.TableNameWriteMarker).
		Where(SQLWhereBetweenSequences, startSequence, endSequence).
		Find(&list).Error

	if err != nil {
		return nil, errors.ThrowLog(err.Error(), constants.ErrBadDatabaseOperation)
	}
	if len(list) == 0 {
		return nil, errors.Throw(constants.ErrEntityNotFound, "write_marker_not_found in range allocation_id: "+allocationID, " start_root: "+startAllocationRoot+" end_root: "+endAllocationRoot)
	}
	return list, nil

}

const (
	SQLWhereAllocationIdAndRootRange = "allocation_id = ? and allocation_root in (?,?)"
	SQLWhereBetweenSequences         = "sequence BETWEEN ? AND ?"
)

// DryRun  Creates a prepared statement when executing any SQL and caches them to speed up future calls
// https://gorm.io/docs/performance.html#Caches-Prepared-Statement
func DryRun(db *gorm.DB) {

	// https://gorm.io/docs/session.html#DryRun
	// Session mode
	tx := db.Session(&gorm.Session{PrepareStmt: true, DryRun: true})

	GetChallengeWriteMarker(context.TODO(), tx, "id", "start_root", "end_root") //nolint
}
