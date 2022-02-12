package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type accountController struct {
	accountService *accountService
}

func NewAccountController(accountService *accountService) *accountController {
	return &accountController{accountService: accountService}
}

func (ctl *accountController) HealthCheck(c *fiber.Ctx) error {
	return WriteResponse(c, http.StatusOK, map[string]string{"status": "ok"})
}

func (ctl *accountController) CreateAccount(c *fiber.Ctx) error {
	body := struct {
		UserID int64 `json:"user_id"`
	}{}
	err := c.BodyParser(&body)
	if err != nil {
		return WriteResponse(c, http.StatusBadRequest, fmt.Sprintf("could not parse body: %s", err))
	}
	if body.UserID == 0 {
		return WriteResponse(c, http.StatusBadRequest, fmt.Sprintf("user_id not provided"))
	}
	account, err := ctl.accountService.CreateByUserID(c.Context(), body.UserID)
	if err != nil {
		return WriteResponse(c, http.StatusBadRequest, fmt.Sprintf("could not create: %s", err))
	}
	return WriteResponse(c, http.StatusOK, account)
}

func (ctl *accountController) TopUpAccount(c *fiber.Ctx) error {
	body := struct {
		Amount float64 `json:"amount"`
	}{}
	err := c.BodyParser(&body)
	if err != nil {
		return WriteResponse(c, http.StatusBadRequest, fmt.Sprintf("could not parse body: %s", err))
	}
	userID, err := GetUserID(c)
	if err != nil {
		return WriteResponse(c, http.StatusForbidden, err.Error())
	}
	account, err := ctl.accountService.TopUp(c.Context(), userID, body.Amount)
	if err != nil {
		return WriteResponse(c, http.StatusBadRequest, err.Error())
	}
	return WriteResponse(c, http.StatusOK, account)
}

func (ctl *accountController) WithdrawAccount(c *fiber.Ctx) error {
	body := struct {
		Amount float64 `json:"amount"`
	}{}
	err := c.BodyParser(&body)
	if err != nil {
		return WriteResponse(c, http.StatusBadRequest, fmt.Sprintf("could not parse body: %s", err))
	}
	userID, err := GetUserID(c)
	if err != nil {
		return WriteResponse(c, http.StatusForbidden, err.Error())
	}
	account, err := ctl.accountService.Withdraw(c.Context(), userID, body.Amount)
	if err != nil {
		return WriteResponse(c, http.StatusBadRequest, err.Error())
	}
	return WriteResponse(c, http.StatusOK, account)
}

func WriteResponse(c *fiber.Ctx, status int, resp interface{}) error {
	c = c.Status(status)
	if resp == nil {
		return nil
	}
	return c.JSON(resp)
}
