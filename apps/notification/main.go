package main

import (
	"context"
	"log"
)

func main() {
	if err := LoadConfig(); err != nil {
		log.Fatalf("could not load config: %s", err)
	}
	log.Println("config loaded successfully")
	notificationRepo := NewNotificationRepository(DBConn)
	notificationService := NewNotificationService(notificationRepo)
	notificationController := NewNotificationController(notificationService)
	messageReceiver, err := NewMessageReceiver(RabbitConn, config.QueueName)
	if err != nil {
		log.Fatalf("could not get NewMessageReceiver: %s", err)
	}
	go func() { log.Fatal(StartServer(config.Port, notificationController)) }()
	go func() { log.Fatal(StartConsumer(context.Background(), messageReceiver, notificationService)) }()
	select {}
}
