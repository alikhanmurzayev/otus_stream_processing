package main

import (
	"log"
)

func main() {
	if err := LoadConfig(); err != nil {
		log.Fatalf("could not load config: %s", err)
	}
	log.Println("config loaded successfully")
	accountRepository := NewAccountRepository(DBConn)
	accountService := NewAccountService(accountRepository)
	accountController := NewAccountController(accountService)
	log.Fatal(StartServer(config.Port, accountController))
}
