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

func (repo *userRepository) CreateUser(ctx context.Context, user User) (User, error) {
	return user, repo.db.WithContext(ctx).Create(&user).Error
}

func (repo *userRepository) GetUser(ctx context.Context, id int64) (User, error) {
	var user User
	return user, repo.db.WithContext(ctx).First(&user, id).Error
}

func (repo *userRepository) UpdateUser(ctx context.Context, user User) (User, error) {
	return user, repo.db.Debug().WithContext(ctx).Save(&user).Error
}
