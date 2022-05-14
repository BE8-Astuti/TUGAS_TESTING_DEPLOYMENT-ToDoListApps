package controller

import (
	"projek/be8/delivery/middlewares"
	"projek/be8/delivery/view"
	vproduk "projek/be8/delivery/view/produk"
	"projek/be8/repository/produk"
	"strconv"

	"net/http"
	"projek/be8/entities"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type ProdukController struct {
	Repo  produk.ProdukRepository
	Valid *validator.Validate
}

func New(NewAddp produk.ProdukRepository, validate *validator.Validate) *ProdukController {
	return &ProdukController{
		Repo:  NewAddp,
		Valid: validate,
	}
}

func (pc *ProdukController) InsertProd() echo.HandlerFunc {
	return func(c echo.Context) error {
		var tmpProd vproduk.InsertProdukRequest

		if err := c.Bind(&tmpProd); err != nil {
			log.Warn("salah input")
			return c.JSON(http.StatusUnsupportedMediaType, view.BindData())
		}

		if err := pc.Valid.Struct(&tmpProd); err != nil {
			log.Warn(err.Error())
			return c.JSON(http.StatusNotAcceptable, view.Validate())
		}
		id := middlewares.ExtractTokenUserId(c)
		newProd := entities.Produk{
			UserID: int(id),
			Nama:   "tango",
			Stok:   10,
		}
		res, errInsert := pc.Repo.InsertProduk(newProd)

		if errInsert != nil {
			log.Warn(errInsert)
			return c.JSON(http.StatusInternalServerError, view.InternalServerError())
		}
		log.Info("berhasil insert")
		return c.JSON(http.StatusCreated, vproduk.StatusCreate(res))

	}
}

func (pc *ProdukController) GetAllProd() echo.HandlerFunc {
	return func(c echo.Context) error {

		res, err := pc.Repo.GetAllProduk()

		if err != nil {
			log.Warn("masalah pada server")
			return c.JSON(http.StatusInternalServerError, view.InternalServerError())
		}
		log.Info("berhasil get all data produk")
		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    http.StatusOK,
			"count":   len(res),
			"message": "berhasil get all data produk",
			"status":  true,
			"data":    res,
		})
	}
}

func (pc *ProdukController) GetProdID() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		convID, err := strconv.Atoi(id)
		if err != nil {
			log.Error(err)
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    http.StatusInternalServerError,
				"message": "connot convert ID",
				"data":    nil,
			})
		}

		hasil, err := pc.Repo.GetProdukID(convID)

		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotFound, view.NotFound())
		}

		return c.JSON(http.StatusOK, vproduk.StatusGetIdOk(hasil))
	}

}
