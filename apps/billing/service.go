package main

import (
	"context"
	"fmt"
)

type accountService struct {
	accountRepo *accountRepository
}

func NewAccountService(accountRepo *accountRepository) *accountService {
	return &accountService{accountRepo: accountRepo}
}

func (service *accountService) GetByUserID(ctx context.Context, userID int64) (Account, error) {
	return service.accountRepo.GetByUserID(ctx, userID)
}

func (service *accountService) CreateByUserID(ctx context.Context, userID int64) (Account, error) {
	account := Account{UserID: userID}
	return service.accountRepo.Save(ctx, account)
}

func (service *accountService) TopUp(ctx context.Context, userID int64, amount float64) (Account, error) {
	account, err := service.accountRepo.GetByUserID(ctx, userID)
	if err != nil {
		return Account{}, fmt.Errorf("GetByUserID: %w", err)
	}
	account.Amount += amount
	return service.accountRepo.Save(ctx, account)
}

func (service *accountService) Withdraw(ctx context.Context, userID int64, amount float64) (Account, error) {
	account, err := service.accountRepo.GetByUserID(ctx, userID)
	if err != nil {
		return Account{}, fmt.Errorf("GetByUserID: %w", err)
	}
	if account.Amount < amount {
		return Account{}, ErrInsufficientBalance
	}
	account.Amount -= amount
	return service.accountRepo.Save(ctx, account)
}
