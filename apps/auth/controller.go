package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
	"strings"
)

type authController struct {
	authService *authService
}

func NewAuthController(authService *authService) *authController {
	return &authController{authService: authService}
}

func (ctl *authController) HealthCheck(c *fiber.Ctx) error {
	return WriteResponse(c, http.StatusOK, map[string]string{"status": "ok"})
}

type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (ctl *authController) Login(c *fiber.Ctx) error {
	var request LoginRequest
	err := c.BodyParser(&request)
	if err != nil {
		return WriteResponse(c, http.StatusBadRequest, fmt.Sprintf("could not parse body: %s", err))
	}
	token, userID, err := ctl.authService.Login(c.Context(), request.Login, request.Password)
	if err != nil {
		return WriteResponse(c, http.StatusUnauthorized, err.Error())
	}
	SetupHeaderUserID(c, userID)
	SetupHeaderAuthToken(c, token)
	return WriteResponse(c, http.StatusOK, nil)
}

func (ctl *authController) IsTokenValid(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	token = strings.TrimPrefix(token, "Bearer")
	token = strings.TrimSpace(token)
	if token == "" {
		return WriteResponse(c, http.StatusUnauthorized, "token not provided")
	}
	valid, userID, err := ctl.authService.IsTokenValid(c.Context(), token)
	if err != nil {
		return WriteResponse(c, http.StatusUnauthorized, err.Error())
	}
	if !valid {
		return WriteResponse(c, http.StatusUnauthorized, nil)
	}
	SetupHeaderUserID(c, userID)
	SetupHeaderAuthToken(c, token)
	return WriteResponse(c, http.StatusOK, nil)
}

func SetupHeaderUserID(c *fiber.Ctx, userID int64) {
	c.Set("x-user-id", strconv.FormatInt(userID, 10))
}

func SetupHeaderAuthToken(c *fiber.Ctx, token string) {
	c.Set("x-auth-token", token)
}

func WriteResponse(c *fiber.Ctx, status int, resp interface{}) error {
	c = c.Status(status)
	if resp == nil {
		return nil
	}
	return c.JSON(resp)
}
