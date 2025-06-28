package repository

import (
	"context"
	"github.com/jmoiron/sqlx"

	"github.com/NewChakrit/golang_gin_vuejs_full-stack/entity"
)

type TransactionRepository interface {
	Add(u entity.Transaction) error
	Edit(ctx context.Context, u entity.Transaction) error
	Delete(id int64) error
	FindAll() ([]entity.Transaction, error)
}

type pgTransactionRepository struct {
	DB *sqlx.DB
}

func NewTransactionRepository(db *sqlx.DB) TransactionRepository {
	return &pgTransactionRepository{db}
}

func (r *pgTransactionRepository) Add(u entity.Transaction) error {
	query := "INSERT INTO transactions (type, ticker, volume, price, date) values ($1, $2, $3, $4, $5) RETURNING id"

	err := r.DB.QueryRow(query, u.Type, u.Ticker, u.Volume, u.Price, u.Date).Scan(&u.ID)

	return err
}

func (r *pgTransactionRepository) Edit(ctx context.Context, u entity.Transaction) error {
	query := "UPDATE transactions SET (type, ticker, volume, price, date) = ($1, $2, $3, $4, $5) WHERE id=$6"

	_, err := r.DB.ExecContext(ctx, query, u.Type, u.Ticker, u.Volume, u.Price, u.Date, u.ID)

	return err
}

func (r *pgTransactionRepository) Delete(id int64) error {
	query := "DELETE FROM transactions WHERE id=$1"

	_, err := r.DB.Exec(query, id)

	return err
}

func (r *pgTransactionRepository) FindAll() ([]entity.Transaction, error) {
	transactions := []entity.Transaction{}

	query := "SELECT * FROM transactions"

	if err := r.DB.Select(&transactions, query); err != nil {
		return nil, err
	}

	return transactions, nil
}
