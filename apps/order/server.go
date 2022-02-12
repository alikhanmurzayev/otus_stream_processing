package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func StartServer(port int, orderController *orderController) error {
	app := fiber.New()
	app.Use(cors.New())

	app.Use(Recover)
	app.Get("/health", orderController.HealthCheck)

	app.Post("/create", orderController.Create)

	return app.Listen(":" + strconv.Itoa(port))
}
