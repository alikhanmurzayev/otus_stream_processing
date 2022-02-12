package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type userController struct {
	userService *userService
}

func NewUserController(userService *userService) *userController {
	return &userController{userService: userService}
}

func (ctl *userController) HealthCheck(c *fiber.Ctx) error {
	return WriteResponse(c, http.StatusOK, map[string]string{"status": "ok"})
}

func (ctl *userController) CreateUser(c *fiber.Ctx) error {
	var user User
	err := c.BodyParser(&user)
	if err != nil {
		return WriteResponse(c, http.StatusBadRequest, fmt.Sprintf("could not parse body: %s", err))
	}
	user, err = ctl.userService.CreateUser(c.Context(), user)
	if err != nil {
		return WriteResponse(c, http.StatusBadRequest, err.Error())
	}
	return WriteResponse(c, http.StatusOK, user)
}

func (ctl *userController) GetUser(c *fiber.Ctx) error {
	userID, err := GetUserID(c)
	if err != nil {
		return WriteResponse(c, http.StatusForbidden, err.Error())
	}
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if id != userID {
		return WriteResponse(c, http.StatusForbidden, nil)
	}
	if err != nil {
		return WriteResponse(c, http.StatusBadRequest, fmt.Sprintf("could not parse id: %s", err))
	}
	user, err := ctl.userService.GetUser(c.Context(), id)
	if err != nil {
		return WriteResponse(c, http.StatusNotFound, err.Error())
	}
	return WriteResponse(c, http.StatusOK, user)
}

func (ctl *userController) UpdateUser(c *fiber.Ctx) error {
	userID, err := GetUserID(c)
	if err != nil {
		return WriteResponse(c, http.StatusForbidden, err.Error())
	}
	var user User
	err = c.BodyParser(&user)
	if err != nil {
		return WriteResponse(c, http.StatusBadRequest, fmt.Sprintf("could not parse body: %s", err))
	}
	user.ID, err = strconv.ParseInt(c.Params("id"), 10, 64)
	if user.ID != userID {
		return WriteResponse(c, http.StatusForbidden, nil)
	}
	if err != nil {
		return WriteResponse(c, http.StatusBadRequest, fmt.Sprintf("could not parse id: %s", err))
	}
	_, err = ctl.userService.UpdateUser(c.Context(), user)
	if err != nil {
		return WriteResponse(c, http.StatusNotFound, err.Error())
	}
	return WriteResponse(c, http.StatusOK, nil)
}

func WriteResponse(c *fiber.Ctx, status int, resp interface{}) error {
	c = c.Status(status)
	if resp == nil {
		return nil
	}
	return c.JSON(resp)
}
