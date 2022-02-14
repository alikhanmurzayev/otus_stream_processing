package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func StartServer(port int, notificationController *notificationController) error {
	app := fiber.New()
	app.Use(cors.New())

	app.Use(Recover)
	app.Get("/health", notificationController.HealthCheck)

	app.Get("/find", notificationController.Find)

	return app.Listen(":" + strconv.Itoa(port))
}
