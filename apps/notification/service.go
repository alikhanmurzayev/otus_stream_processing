package main

import (
	"context"
	"fmt"
)

type orderService struct {
	orderRepository  *orderRepository
	billingAdapter   *billingAdapter
	messagePublisher *messagePublisher
}

func NewOrderService(orderRepository *orderRepository, billingAdapter *billingAdapter, messagePublisher *messagePublisher) *orderService {
	return &orderService{
		orderRepository:  orderRepository,
		billingAdapter:   billingAdapter,
		messagePublisher: messagePublisher,
	}
}

func (service *orderService) Create(ctx context.Context, userID int64, price float64) (Order, error) {
	order := Order{
		UserID: userID,
		Price:  price,
		Status: OrderStatusNew,
	}
	order, err := service.orderRepository.Save(ctx, order)
	if err != nil {
		return Order{}, fmt.Errorf("could not create order: %w", err)
	}
	err = service.billingAdapter.WithdrawAccount(ctx, userID, price)
	if err != nil {
		order.Status = OrderStatusUnpaid
	} else {
		order.Status = OrderStatusPaid
	}
	order, err = service.orderRepository.Save(ctx, order)
	if err != nil {
		return Order{}, fmt.Errorf("could not Save: %w", err)
	}
	err = service.messagePublisher.Publish(ctx, OrderEvent{OrderID: order.ID, UserID: userID, Status: order.Status})
	if err != nil {
		return Order{}, fmt.Errorf("could not Publish: %w", err)
	}
	return order, nil
}
