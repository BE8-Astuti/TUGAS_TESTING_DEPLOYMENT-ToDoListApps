package controller

import (
	"projek/be8/delivery/middlewares"
	"projek/be8/delivery/view"
	vtransaksi "projek/be8/delivery/view/transaksi"
	"projek/be8/repository/produk"
	"projek/be8/repository/transaksi"

	"net/http"
	"projek/be8/entities"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type TransaksiController struct {
	Rproduk produk.ProdukRepository
	Repo    transaksi.Transaksi
	Valid   *validator.Validate
}

func New(repo transaksi.Transaksi, rproduk produk.ProdukRepository, valid *validator.Validate) *TransaksiController {
	return &TransaksiController{
		Rproduk: rproduk,
		Repo:    repo,
		Valid:   valid,
	}
}

func (tc *TransaksiController) InsertTransaksi() echo.HandlerFunc {
	return func(c echo.Context) error {
		var tmpTransaksi vtransaksi.InsertTransaksiRequest

		if err := c.Bind(&tmpTransaksi); err != nil {
			log.Warn("salah input")
			return c.JSON(http.StatusUnsupportedMediaType, view.BindData())
		}

		if err := tc.Valid.Struct(tmpTransaksi); err != nil {
			log.Warn(err.Error())
			return c.JSON(http.StatusNotAcceptable, view.Validate())
		}

		// if tmpTransaksi.Jenis_transaksi != "pembelian" && tmpTransaksi.Jenis_transaksi != "penjualan" {
		// 	return c.JSON(http.StatusBadRequest, "fail")
		// }

		newTransaksi := entities.Transaksi{
			User_id:         tmpTransaksi.User_id,
			Produk_id:       tmpTransaksi.Produk_id,
			Produk:          tmpTransaksi.Produk,
			Qty:             tmpTransaksi.Qty,
			Jenis_transaksi: tmpTransaksi.Jenis_transaksi}
		res, err := tc.Repo.Insert(newTransaksi)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusInternalServerError, view.InternalServerError())
		}

		produk, _ := tc.Rproduk.GetProdukID(int(res.Produk_id))

		if newTransaksi.Jenis_transaksi == "pembelian" {
			produk.Stok += tmpTransaksi.Qty
		}

		if newTransaksi.Jenis_transaksi == "penjualan" {
			produk.Stok -= tmpTransaksi.Qty
		}
		user_id := middlewares.ExtractTokenUserId(c)
		_, err = tc.Rproduk.UpdateProduk(int(user_id), produk, int(res.Produk_id))

		if err != nil {
			log.Warn("tidak bisa update stok")
			return c.JSON(http.StatusInternalServerError, view.InternalServerError())
		}

		log.Info("berhasil insert")
		return c.JSON(http.StatusCreated, vtransaksi.StatusCreate(res))
	}
}

func (tc *TransaksiController) RiwayatAllTrans() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := tc.Repo.RiwayatAllTrans()

		if err != nil {
			log.Warn("masalah pada server")
			return c.JSON(http.StatusInternalServerError, view.InternalServerError())
		}
		log.Info("berhasil get all data")
		return c.JSON(http.StatusOK, vtransaksi.StatusGetAllOk(res))
	}
}

func (tc *TransaksiController) HistoriTrans() echo.HandlerFunc {
	return func(c echo.Context) error {
		tipe := c.Param("tipe")

		hasil, err := tc.Repo.HistoriTrans(tipe)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusInternalServerError, view.InternalServerError())
		}

		log.Info("data found")
		return c.JSON(http.StatusOK, vtransaksi.StatusGetAllOk(hasil))
	}

}

func (tc *TransaksiController) GetAllTransaksi() echo.HandlerFunc {
	return func(c echo.Context) error {

		res, err := tc.Repo.GetAll()

		if err != nil {
			log.Warn("masalah pada server")
			return c.JSON(http.StatusInternalServerError, view.InternalServerError())
		}
		log.Info("berhasil get all data")
		return c.JSON(http.StatusOK, vtransaksi.StatusGetAllOk(res))

	}
}

func (tc *TransaksiController) GetTransaksi() echo.HandlerFunc {
	return func(c echo.Context) error {
		tipe := c.Param("tipe")

		hasil, err := tc.Repo.GetTrans(tipe)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusInternalServerError, view.InternalServerError())
		}

		log.Info("data found")
		return c.JSON(http.StatusOK, vtransaksi.StatusGetAllOk(hasil))
	}

}
