package main

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/google/uuid"
)

type userService struct {
	userRepository *userRepository
	billingAdapter *billingAdapter
}

func NewUserService(userRepository *userRepository, billingAdapter *billingAdapter) *userService {
	return &userService{
		userRepository: userRepository,
		billingAdapter: billingAdapter,
	}
}

func (service *userService) CreateUser(ctx context.Context, user User) (User, error) {
	user.Password, user.Salt = service.hashWithSalt(user.Password)
	user, err := service.userRepository.CreateUser(ctx, user)
	if err != nil {
		return User{}, fmt.Errorf("CreateUser: %w", err)
	}
	return user, service.billingAdapter.CreateAccount(ctx, user.ID)
}

func (service *userService) GetUser(ctx context.Context, id int64) (User, error) {
	return service.userRepository.GetUser(ctx, id)
}

func (service *userService) UpdateUser(ctx context.Context, user User) (User, error) {
	return service.userRepository.UpdateUser(ctx, user)
}

func (service *userService) hashWithSalt(password string) (hashedPassword, salt string) {
	salt = uuid.New().String()
	hashedArray := sha256.Sum256([]byte(password + salt))
	hashedPassword = hex.EncodeToString(hashedArray[:])
	return
}
