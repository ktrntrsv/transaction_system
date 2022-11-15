package db

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/ktrntrsv/transactionService/internal/domain"
	"github.com/ktrntrsv/transactionService/pkg/logger"
)

type AccountRepository struct {
	*Database
	logger logger.Interface
}

func NewAccountRepository(client *Database, logger logger.Interface) *AccountRepository {
	return &AccountRepository{
		Database: client,
		logger:   logger,
	}
}

func (r *AccountRepository) GetByIdWithLock(ctx context.Context, id uuid.UUID) (domain.AccountModel, error) {
	query := `SELECT id, balance from account WHERE id = $1 FOR UPDATE;`
	var account domain.AccountModel
	row := r.model(ctx).QueryRow(ctx, query, id)
	err := row.Scan(&account.Id, &account.Balance)
	if err != nil {
		if err == pgx.ErrNoRows {
			return domain.AccountModel{}, domain.ErrAccountNotFound
		}
		return domain.AccountModel{}, err
	}
	return account, nil
}

func (r *AccountRepository) UpdateBalance(ctx context.Context, account domain.AccountModel) error {
	query := `UPDATE account SET balance=$1 WHERE id=$2;`
	_, err := r.model(ctx).Exec(ctx, query, account.Balance, account.Id)
	if err != nil {
		return err
	}
	return nil
}
