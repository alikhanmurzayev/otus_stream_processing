package main

import (
	"context"
	"gorm.io/gorm"
)

type accountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) *accountRepository {
	return &accountRepository{db: db}
}

func (repo *accountRepository) GetByUserID(ctx context.Context, userID int64) (Account, error) {
	account := Account{UserID: userID}
	return account, repo.db.WithContext(ctx).Where(&account).First(&account).Error
}

func (repo *accountRepository) Save(ctx context.Context, account Account) (Account, error) {
	return account, repo.db.WithContext(ctx).Save(&account).Error
}
