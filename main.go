package main

import (
	
	"projek/be8/config"
	cproduk "projek/be8/delivery/controller/produk"
	ctransaksi "projek/be8/delivery/controller/transaksi"
	cuser "projek/be8/delivery/controller/user"
	"projek/be8/utils"

	// cbook "mware/be8/delivery/controller/book"
	"projek/be8/delivery/routes"

	produkRepo "projek/be8/repository/produk"
	transaksiRepo "projek/be8/repository/transaksi"
	userRepo "projek/be8/repository/user"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func main() {
	conf := config.GetConfig()
	db := utils.InitDB(conf)
	// utils.InitMigrate(db)
	e := echo.New()

	repoUser := userRepo.New(db)
	repoTransaksi := transaksiRepo.New(db)
	repoProduk := produkRepo.New(db)

	controllUser := cuser.New(repoUser, validator.New())
	controllTransaksi := ctransaksi.New(repoTransaksi, repoProduk, validator.New())
	controllProduk := cproduk.New(repoProduk, validator.New())

	routes.RegisterPath(e, controllProduk, controllUser, controllTransaksi)

	e.Logger.Fatal(e.Start(":8000"))
}
