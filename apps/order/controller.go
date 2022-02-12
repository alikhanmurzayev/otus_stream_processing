package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type orderController struct {
	orderService *orderService
}

func NewOrderController(orderService *orderService) *orderController {
	return &orderController{orderService: orderService}
}

func (ctl *orderController) HealthCheck(c *fiber.Ctx) error {
	return WriteResponse(c, http.StatusOK, map[string]string{"status": "ok"})
}

func (ctl *orderController) Create(c *fiber.Ctx) error {
	body := struct {
		Price float64 `json:"price"`
	}{}
	err := c.BodyParser(&body)
	if err != nil {
		return WriteResponse(c, http.StatusBadRequest, fmt.Sprintf("could not parse body: %s", err))
	}
	userID, err := GetUserID(c)
	if err != nil {
		return WriteResponse(c, http.StatusForbidden, err.Error())
	}
	order, err := ctl.orderService.Create(c.Context(), userID, body.Price)
	if err != nil {
		return WriteResponse(c, http.StatusBadRequest, err.Error())
	}
	return WriteResponse(c, http.StatusOK, order)
}

func WriteResponse(c *fiber.Ctx, status int, resp interface{}) error {
	c = c.Status(status)
	if resp == nil {
		return nil
	}
	return c.JSON(resp)
}
