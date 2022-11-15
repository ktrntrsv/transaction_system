package domain

import (
	"context"
	"errors"
	"github.com/google/uuid"
)

var ErrAccountNotFound = errors.New("account not found")

type accountRepository interface {
	GetByIdWithLock(ctx context.Context, id uuid.UUID) (AccountModel, error)
	UpdateBalance(ctx context.Context, account AccountModel) error
	WithinTransaction(ctx context.Context, tFunc func(ctx context.Context) error) error
}

type AccountUsecase struct {
	accRepo accountRepository
}

func NewAccountUsecase(accRepo accountRepository) *AccountUsecase {
	return &AccountUsecase{accRepo: accRepo}
}

func (uc *AccountUsecase) TransferMoney(ctx context.Context, accFromId uuid.UUID, accToId uuid.UUID, amount float64) error {

	err := uc.accRepo.WithinTransaction(
		ctx,
		func(ctx context.Context) error {
			accFrom, err := uc.accRepo.GetByIdWithLock(ctx, accFromId)
			if err != nil {
				return ErrAccountNotFound
			}
			accTo, err := uc.accRepo.GetByIdWithLock(ctx, accToId)
			if err != nil {
				return ErrAccountNotFound
			}

			// acc from balance updating
			err = accFrom.UpdateBalance(-amount)
			if err != nil {
				return err
			}
			err = uc.accRepo.UpdateBalance(ctx, accFrom)
			if err != nil {
				return err
			}

			// acc to balance updating
			err = accTo.UpdateBalance(amount)
			if err != nil {
				return err
			}
			err = uc.accRepo.UpdateBalance(ctx, accTo)
			if err != nil {
				return err
			}

			return nil
		})
	return err
}
