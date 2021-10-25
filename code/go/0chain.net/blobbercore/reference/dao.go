package reference

import (
	"context"

	"github.com/0chain/blobber/code/go/0chain.net/blobbercore/models"
	"github.com/0chain/errors"
	"github.com/0chain/gosdk/constants"
	"gorm.io/gorm"
)

func Exists(ctx context.Context, db *gorm.DB, allocationTx, path string) (bool, error) {
	if len(allocationTx) == 0 {
		return false, errors.Throw(constants.ErrInvalidParameter, "allocationTx")
	}

	if len(path) == 0 {
		return false, errors.Throw(constants.ErrInvalidParameter, "path")
	}

	var count int64
	err := db.Table(models.TableNameReferenceObject).Where(SQLWhereAllocationTxAndPath, allocationTx, path).Count(&count).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}

		return false, errors.ThrowLog(err.Error(), constants.ErrBadDatabaseOperation)
	}

	return count > 0, nil
}

func Get(ctx context.Context, db *gorm.DB, allocationTx, path string) (*models.ReferenceObject, error) {

	if len(allocationTx) == 0 {
		return nil, errors.Throw(constants.ErrInvalidParameter, "allocationTx")
	}

	if len(path) == 0 {
		return nil, errors.Throw(constants.ErrInvalidParameter, "path")
	}

	it := &models.ReferenceObject{}

	result := db.Table(models.TableNameReferenceObject).Where(SQLWhereAllocationTxAndPath, allocationTx, path).First(it)

	if result.Error == nil {
		return it, nil
	}

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return it, errors.Throw(constants.ErrEntityNotFound, "allocation_id: "+allocationTx+" path: "+path)
	}

	return nil, errors.ThrowLog(result.Error.Error(), constants.ErrBadDatabaseOperation)
}

// GetChallengeMeta get total blocks number by allocation_tx
func GetChallengeMeta(ctx context.Context, db *gorm.DB, allocationTx string) (*models.ChallengeMeta, error) {
	if len(allocationTx) == 0 {
		return nil, errors.Throw(constants.ErrInvalidParameter, "allocationTx")
	}

	it := &models.ChallengeMeta{}

	err := db.Table(models.TableNameReferenceObject).
		Where(SQLWhereAllocationTxAndPath, allocationTx, "/").
		First(it).Error

	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.Throw(constants.ErrEntityNotFound, "allocation_id: "+allocationTx+" path: /")
		}

		return nil, errors.ThrowLog(err.Error(), constants.ErrBadDatabaseOperation)
	}

	return it, nil
}

// GetNextBlocks get blocks object for challenge by allocation_id and parentPath
func GetNextBlocks(ctx context.Context, db *gorm.DB, allocationTx string, parentPath string) ([]*models.StorageObject, error) {

	if len(allocationTx) == 0 {
		return nil, errors.Throw(constants.ErrInvalidParameter, "allocationTx")
	}

	if len(parentPath) == 0 {
		return nil, errors.Throw(constants.ErrInvalidParameter, "parentPath")
	}

	var list []*models.StorageObject

	err := db.Table(models.TableNameReferenceObject).
		Where(SQLWhereAllocationTxAndParentPath, allocationTx, parentPath).
		Order("level, type desc, lookup_hash"). // first files, and dirs
		Find(&list).Error

	if err != nil {
		return nil, errors.ThrowLog(err.Error(), constants.ErrBadDatabaseOperation)
	}

	return list, nil
}

const (
	SQLWhereAllocationTxAndPath       = "reference_objects.allocation_id = ? and reference_objects.path = ? and deleted_at is NULL"
	SQLWhereAllocationTxAndParentPath = "reference_objects.allocation_id = ? and reference_objects.parent_path = ? and deleted_at is NULL"
)

// DryRun  Creates a prepared statement when executing any SQL and caches them to speed up future calls
// https://gorm.io/docs/performance.html#Caches-Prepared-Statement
func DryRun(db *gorm.DB) {

	// https://gorm.io/docs/session.html#DryRun
	// Session mode
	tx := db.Session(&gorm.Session{PrepareStmt: true, DryRun: true})

	// use Table instead of Model to reduce reflect times

	// prepare statement for GetOrCreate
	tx.Table(models.TableNameAllocation).Where(SQLWhereAllocationTxAndPath, "allocation_id", "path").First(&models.ReferenceObject{})
	// count
	var count int64
	tx.Table(models.TableNameAllocation).Where(SQLWhereAllocationTxAndPath, "allocation_id", "path").Count(&count)

}
