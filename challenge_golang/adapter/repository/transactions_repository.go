package repository

import (
	"database/sql"
	"time"
)

type TransactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (t *TransactionRepository) Insert(id string, accountId string, amount float64) error {
	stmt, err := t.db.Prepare(`
		Insert into transactions (id, account_id, amount, created_at)
		values($1, $2, $3, $4)
	`)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		id,
		accountId,
		amount,
		time.Now(),
	)

	if err != nil {
		return err
	}

	return nil
}
