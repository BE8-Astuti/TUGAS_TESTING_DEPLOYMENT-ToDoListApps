package controller

import (
	"github.com/labstack/echo/v4"
)

type ControllerUser interface {
	InsertUser() echo.HandlerFunc
	Login() echo.HandlerFunc
}
