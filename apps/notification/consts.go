package main

type OrderStatus string

const (
	OrderStatusNew    OrderStatus = "new"
	OrderStatusPaid   OrderStatus = "paid"
	OrderStatusUnpaid OrderStatus = "unpaid"
)
