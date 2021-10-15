package allocation

import (
	"context"

	"github.com/0chain/blobber/code/go/0chain.net/blobbercore/datastore"
	"github.com/0chain/blobber/code/go/0chain.net/blobbercore/models"
	"github.com/0chain/errors"
	"github.com/0chain/gosdk/constants"
	"gorm.io/gorm"
)

// GetOrCreate, get allocation if it exists in db. if not, try to sync it from blockchain, and insert it in db.
func GetOrCreate(ctx context.Context, store datastore.Store, allocationTx string) (*models.Allocation, error) {
	db := store.GetDB()

	if len(allocationTx) == 0 {
		return nil, errors.Throw(constants.ErrInvalidParameter, "tx")
	}

	alloc := &models.Allocation{}
	result := db.Table(models.TableNameAllocation).Where(SQLWhereGetByTx, allocationTx).First(alloc)

	if result.Error == nil {
		return alloc, nil
	}

	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.ThrowLog(result.Error.Error(), constants.ErrBadDatabaseOperation)
	}

	return SyncAllocation(allocationTx)

}

const (
	SQLWhereGetByTx = "allocations.tx = ?"
)

// DryRun  Creates a prepared statement when executing any SQL and caches them to speed up future calls
// https://gorm.io/docs/performance.html#Caches-Prepared-Statement
func DryRun(db *gorm.DB) {

	// Session mode
	tx := db.Session(&gorm.Session{PrepareStmt: true})

	// use Table instead of Model to reduce reflect times

	// prepare statement for GetOrCreate
	tx.Table(models.TableNameAllocation).Where(SQLWhereGetByTx, "tx").First(&models.Allocation{})

}
