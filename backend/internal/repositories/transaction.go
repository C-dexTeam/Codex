package repositories

import (
	"github.com/C-dexTeam/codex/internal/domains"

	"github.com/jmoiron/sqlx"
)

type TransactionRepository struct {
	db *sqlx.DB
}

func NewTransactionRepository(db *sqlx.DB) domains.ITransactionRepository {
	return &TransactionRepository{db: db}
}

// Begin starts a new transaction
func (tr *TransactionRepository) Begin() (*sqlx.Tx, error) {
	return tr.db.Beginx()
}

// Commit commits the current transaction
func (tr *TransactionRepository) Commit(tx *sqlx.Tx) error {
	return tx.Commit()
}
