package main

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type ctxError string

func (e ctxError) Error() string {
	return string(e)
}

const (
	ErrUserIDNotProvided   = ctxError("user id not provided")
	ErrUserIDInvalidFormat = ctxError("invalid format of user id")
)

func GetUserID(c *fiber.Ctx) (int64, error) {
	rawUserID := c.Request().Header.Peek("x-user-id")
	if len(rawUserID) == 0 {
		return 0, ErrUserIDNotProvided
	}
	userID, err := strconv.ParseInt(string(rawUserID), 10, 64)
	if err != nil {
		return 0, ErrUserIDInvalidFormat
	}
	return userID, nil

}
