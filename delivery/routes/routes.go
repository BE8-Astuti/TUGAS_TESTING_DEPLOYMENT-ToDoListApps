package routes

import (
	cproduk "projek/be8/delivery/controller/produk"
	ctransaksi "projek/be8/delivery/controller/transaksi"
	cuser "projek/be8/delivery/controller/user"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterPath(e *echo.Echo, pc cproduk.ControllerProduk, uc cuser.ControllerUser, tc ctransaksi.ControllerTransaksi) {
	// e.Pre(middleware.AddTrailingSlash())
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time:${time_rfc3339}, method=${method}, uri=${uri}, status=${status}\n",
	}))

	e.POST("/user", uc.InsertUser()) // Register
	e.POST("/login", uc.Login())     // Login

	e.POST("/produk", pc.InsertProd(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("A$T0T!")}))
	e.GET("/produk", pc.GetAllProd(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("A$T0T!")}))
	e.GET("/produk/:id", pc.GetProdID(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("A$T0T!")}))

	e.POST("/transaksi", tc.InsertTransaksi(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("A$T0T!")}))
	e.GET("/transaksi", tc.GetAllTransaksi(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("A$T0T!")}))
	e.GET("/transaksi/:tipe", tc.GetTransaksi(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("A$T0T!")}))
	e.GET("/transaksi/recently7days", tc.RiwayatAllTrans(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("A$T0T!")}))
	e.GET("/transaksi/:tipe/last-7-days", tc.HistoriTrans(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("A$T0T!")}))

}
