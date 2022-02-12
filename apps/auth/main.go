package main

import (
	"log"
)

func main() {
	if err := LoadConfig(); err != nil {
		log.Fatalf("could not load config: %s", err)
	}
	userRepository := NewUserRepository(DBConn)
	authService := NewAuthService(userRepository)
	authController := NewAuthController(authService)
	log.Fatal(StartServer(config.Port, authController))
}
