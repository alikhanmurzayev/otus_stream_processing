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
	messagePublisher, err := NewMessagePublisher(RabbitConn, config.QueueName)
	if err != nil {
		log.Fatalf("could not get NewMessagePublisher: %s", err)
	}
	orderService := NewOrderService(orderRepo, billingAdapter, messagePublisher)
	orderController := NewOrderController(orderService)
	log.Fatal(StartServer(config.Port, orderController))
}
