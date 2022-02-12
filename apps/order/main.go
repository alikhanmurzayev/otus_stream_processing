package main

import (
	"log"
)

func main() {
	if err := LoadConfig(); err != nil {
		log.Fatalf("could not load config: %s", err)
	}
	log.Println("config loaded successfully")
	orderRepo := NewOrderRepository(DBConn)
	billingAdapter := NewBillingAdapter(config.BillingService)
	orderService := NewOrderService(orderRepo, billingAdapter)
	orderController := NewOrderController(orderService)
	log.Fatal(StartServer(config.Port, orderController))
}
