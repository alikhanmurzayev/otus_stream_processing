package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func StartServer(port int, authController *authController) error {
	app := fiber.New()
	app.Use(cors.New())

	app.Use(Recover)
	app.Get("/health", authController.HealthCheck)

	app.Post("/login", authController.Login)
	app.All("/is-token-valid*", authController.IsTokenValid)

	return app.Listen(":" + strconv.Itoa(port))
}
