package controllers

import (
	"go-account-api/core"
	"go-account-api/schemas"

	"github.com/labstack/echo"
)

func Login(c echo.Context) error {
	req := new(schemas.LoginRequest)

	if err := c.Bind(req); err != nil {
		return echo.ErrBadRequest
	}

	res, status := core.Login(req)
	return c.JSON(status, res)
}
func Register(c echo.Context) error {
	req := new(schemas.RegisterRequest)

	if err := c.Bind(req); err != nil {
		return echo.ErrBadRequest
	}

	res, status := core.Register(req)
	return c.JSON(status, res)
}

func Logout(c echo.Context) error {
	req := new(schemas.LogoutRequest)

	if err := c.Bind(req); err != nil {
		return echo.ErrBadRequest
	}

	res, status := core.Logout(req)
	return c.JSON(status, res)
}
