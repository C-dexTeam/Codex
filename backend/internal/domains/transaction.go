package domains

import (
	"github.com/jmoiron/sqlx"
)

// ITransactionRepository is the interface that provides the methods for the Transaction repository.
type ITransactionRepository interface {
	Begin() (*sqlx.Tx, error)
	Commit(tx *sqlx.Tx) error
}
