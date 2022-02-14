package main

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type notificationRepository struct {
	db *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) *notificationRepository {
	return &notificationRepository{db: db}
}

func (repo *notificationRepository) Save(ctx context.Context, notification Notification) (Notification, error) {
	return notification, repo.db.WithContext(ctx).Save(&notification).Error
}

func (repo *notificationRepository) Find(ctx context.Context, userID int64, page, size int) (notifications []Notification, total int64, err error) {
	query := repo.db.WithContext(ctx).Model(&Notification{}).Where(&Notification{UserID: userID})
	err = query.Count(&total).Error
	if err != nil {
		err = fmt.Errorf("count error: %w", err)
		return
	}
	err = query.Order("id DESC").Offset(page * size).Limit(size).Find(&notifications).Error
	return
}
