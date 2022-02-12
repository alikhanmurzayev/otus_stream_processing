package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func StartServer(port int, userController *userController) error {
	app := fiber.New()
	app.Use(cors.New())

	app.Use(Recover)
	app.Get("/health", userController.HealthCheck)

	app.Post("/register", userController.CreateUser)
	app.Get("/:id", userController.GetUser)
	app.Put("/:id", userController.UpdateUser)

	return app.Listen(":" + strconv.Itoa(port))
}
