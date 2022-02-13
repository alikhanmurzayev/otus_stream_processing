package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func Recover(c *fiber.Ctx) error {
	defer func() {
		if r := recover(); r != nil {
			c.Status(http.StatusInternalServerError)
			_, _ = c.WriteString(fmt.Sprintf("%s", r))
		}
	}()
	return c.Next()
}
