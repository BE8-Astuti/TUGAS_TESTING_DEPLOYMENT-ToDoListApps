package controller

import (
	"github.com/labstack/echo/v4"
)

type ControllerTransaksi interface {
	InsertTransaksi() echo.HandlerFunc

	GetAllTransaksi() echo.HandlerFunc
	GetTransaksi() echo.HandlerFunc
	RiwayatAllTrans() echo.HandlerFunc
	HistoriTrans() echo.HandlerFunc
}
