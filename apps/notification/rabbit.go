package main

import (
	"context"
	"fmt"
	"log"
)

func StartConsumer(ctx context.Context, receiver *messageReceiver, notificationService *notificationService) error {
	messages, err := receiver.Receive(ctx)
	if err != nil {
		return fmt.Errorf("could not Receive: %w", err)
	}
	for orderEvent := range messages {
		notification := GetNotificationFromOrderEvent(orderEvent)
		err := notificationService.Send(context.Background(), notification)
		if err != nil {
			log.Printf("could not send %#v: %s", notification, err)
		} else {
			log.Printf("successfully sent %#v", notification)
		}
	}
	return nil
}

func GetNotificationFromOrderEvent(orderEvent OrderEvent) Notification {
	notification := Notification{UserID: orderEvent.UserID}
	switch orderEvent.Status {
	case OrderStatusPaid:
		notification.Text = fmt.Sprintf("happiness letter for order with id %d", orderEvent.OrderID)
	case OrderStatusUnpaid:
		notification.Text = fmt.Sprintf("grief letter for order with id %d", orderEvent.OrderID)
	case OrderStatusNew:
		notification.Text = fmt.Sprintf("new order with id %d", orderEvent.OrderID)
	}
	return notification
}
