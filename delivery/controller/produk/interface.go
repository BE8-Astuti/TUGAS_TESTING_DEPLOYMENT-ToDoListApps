package controller

import (
	"github.com/labstack/echo/v4"
)

type ControllerProduk interface {
	InsertProd() echo.HandlerFunc
	GetAllProd() echo.HandlerFunc
	GetProdID() echo.HandlerFunc
}
