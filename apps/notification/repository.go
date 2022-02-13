package main

import (
	"context"
	"gorm.io/gorm"
)

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *orderRepository {
	return &orderRepository{db: db}
}

func (repo *orderRepository) Save(ctx context.Context, order Order) (Order, error) {
	return order, repo.db.WithContext(ctx).Save(&order).Error
}

func (repo *orderRepository) GetByID(ctx context.Context, id int64) (Order, error) {
	var order Order
	return order, repo.db.WithContext(ctx).Where(&Order{ID: id}).First(&order).Error
}
