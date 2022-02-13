package main

type OrderEvent struct {
	OrderID int64       `json:"order_id"`
	UserID  int64       `json:"user_id"`
	Status  OrderStatus `json:"status"`
}
