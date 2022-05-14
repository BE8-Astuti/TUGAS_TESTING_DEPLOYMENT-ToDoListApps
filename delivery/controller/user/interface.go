package controller

import (
	"github.com/labstack/echo/v4"
)

type ControllerUser interface {
	InsertUser(c echo.Context) error
	Login(c echo.Context) error
}
