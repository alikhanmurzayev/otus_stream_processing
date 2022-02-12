package main

import (
	"context"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (repo *userRepository) GetByLogin(ctx context.Context, login string) (User, error) {
	var user User
	return user, repo.db.WithContext(ctx).Where(&User{Login: login}).First(&user).Error
}
