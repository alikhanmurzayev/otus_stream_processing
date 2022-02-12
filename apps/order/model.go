package main

type Order struct {
	ID     int64       `json:"id"`
	UserID int64       `json:"user_id"`
	Price  float64     `json:"price"`
	Status OrderStatus `json:"status"`
}

func (*Order) TableName() string {
	return "orders"
}
