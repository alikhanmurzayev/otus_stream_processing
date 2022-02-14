package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
)

type notificationController struct {
	notificationService *notificationService
}

func NewNotificationController(notificationService *notificationService) *notificationController {
	return &notificationController{notificationService: notificationService}
}

func (ctl *notificationController) HealthCheck(c *fiber.Ctx) error {
	return WriteResponse(c, http.StatusOK, map[string]string{"status": "ok"})
}

func (ctl *notificationController) Find(c *fiber.Ctx) error {
	userID, err := GetUserID(c)
	if err != nil {
		return WriteResponse(c, http.StatusForbidden, err.Error())
	}
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		return WriteResponse(c, http.StatusBadRequest, fmt.Sprintf("could not parse page: %s", err))
	}
	size, err := strconv.Atoi(c.Query("size"))
	if err != nil {
		return WriteResponse(c, http.StatusBadRequest, fmt.Sprintf("could not parse size: %s", err))
	}
	resp := struct {
		Rows  []Notification `json:"rows"`
		Total int64          `json:"total"`
	}{}
	resp.Rows, resp.Total, err = ctl.notificationService.Find(c.Context(), userID, page, size)
	if err != nil {
		return WriteResponse(c, http.StatusBadRequest, err.Error())
	}
	return WriteResponse(c, http.StatusOK, resp)
}

func WriteResponse(c *fiber.Ctx, status int, resp interface{}) error {
	c = c.Status(status)
	if resp == nil {
		return nil
	}
	return c.JSON(resp)
}
