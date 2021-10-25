package allocation

import (
	"context"

	"github.com/0chain/blobber/code/go/0chain.net/blobbercore/models"
	"github.com/0chain/errors"
	"github.com/0chain/gosdk/constants"
	"gorm.io/gorm"
)

// GetOrCreate, get allocation if it exists in db. if not, try to sync it from blockchain, and insert it in db.
func GetOrCreate(ctx context.Context, db *gorm.DB, allocationTx string) (*models.Allocation, error) {

	if len(allocationTx) == 0 {
		return nil, errors.Throw(constants.ErrInvalidParameter, "allocationTx")
	}

	alloc := &models.Allocation{}
	err := db.Table(models.TableNameAllocation).
		Where(SQLWhereTx, allocationTx).
		First(alloc).
		Error

	if err == nil {
		return alloc, nil
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.ThrowLog(err.Error(), constants.ErrBadDatabaseOperation)
	}

	return SyncAllocation(allocationTx)

}

// Get get allocation by tx/id
func Get(ctx context.Context, db *gorm.DB, allocationTx string) (*models.Allocation, error) {
	if len(allocationTx) == 0 {
		return nil, errors.Throw(constants.ErrInvalidParameter, "allocationTx")
	}

	alloc := &models.Allocation{}
	err := db.Table(models.TableNameAllocation).
		Where(SQLWhereTx, allocationTx).
		First(alloc).
		Error

	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.Throw(constants.ErrEntityNotFound, "tx")
		}

		return nil, errors.ThrowLog(err.Error(), constants.ErrBadDatabaseOperation)

	}

	return alloc, nil
}

// GetAllocationRoot get allocation_root by id/tx
func GetAllocationRoot(ctx context.Context, db *gorm.DB, allocationTx string) (string, error) {
	if len(allocationTx) == 0 {
		return "", errors.Throw(constants.ErrInvalidParameter, "allocationTx")
	}

	var result map[string]string
	err := db.Table(models.TableNameAllocation).
		Select("allocation_root").
		Where(SQLWhereTx, allocationTx).
		First(result).
		Error

	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.Throw(constants.ErrEntityNotFound, "tx")
		}

		return "", errors.ThrowLog(err.Error(), constants.ErrBadDatabaseOperation)

	}

	return result["allocation_root"], nil
}

const (
	SQLWhereTx = "allocations.tx = ?"
)

// DryRun  Creates a prepared statement when executing any SQL and caches them to speed up future calls
// https://gorm.io/docs/performance.html#Caches-Prepared-Statement
func DryRun(db *gorm.DB) {

	// https://gorm.io/docs/session.html#DryRun
	// Session mode
	tx := db.Session(&gorm.Session{PrepareStmt: true, DryRun: true})

	// use Table instead of Model to reduce reflect times

	// prepare statement for Get
	var result map[string]string

	//Get
	tx.Table(models.TableNameAllocation).Where(SQLWhereTx, "tx").First(&models.Allocation{})
	//GetAllocationRoot
	tx.Table(models.TableNameAllocation).Select("allocation_root").Where(SQLWhereTx, "tx").First(&result)

}
