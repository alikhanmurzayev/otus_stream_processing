package main

import "context"

type notificationService struct {
	notificationRepository *notificationRepository
}

func NewNotificationService(notificationRepository *notificationRepository) *notificationService {
	return &notificationService{notificationRepository: notificationRepository}
}

func (service *notificationService) Send(ctx context.Context, notification Notification) error {
	_, err := service.notificationRepository.Save(ctx, notification)
	return err
}

func (service *notificationService) Find(ctx context.Context, userID int64, page, size int) (notifications []Notification, total int64, err error) {
	return service.notificationRepository.Find(ctx, userID, page, size)
}
