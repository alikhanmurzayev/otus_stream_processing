package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func StartServer(port int, accountController *accountController) error {
	app := fiber.New()
	app.Use(cors.New())

	app.Use(Recover)
	app.Get("/health", accountController.HealthCheck)

	app.Get("/account", accountController.GetAccount)
	app.Post("/create", accountController.CreateAccount)
	app.Post("/top-up", accountController.TopUpAccount)
	app.Post("/withdraw", accountController.WithdrawAccount)

	return app.Listen(":" + strconv.Itoa(port))
}
