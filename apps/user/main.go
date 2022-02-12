package main

import (
	"log"
)

func main() {
	if err := LoadConfig(); err != nil {
		log.Fatalf("could not load config: %s", err)
	}
	log.Println("config loaded successfully")
	userRepository := NewUserRepository(DBConn)
	billingAdapter := NewBillingAdapter(config.BillingService)
	userService := NewUserService(userRepository, billingAdapter)
	userController := NewUserController(userService)
	log.Fatal(StartServer(config.Port, userController))
}
