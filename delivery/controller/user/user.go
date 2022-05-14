package controller

import (
	"projek/be8/delivery/middlewares"
	"projek/be8/delivery/view"
	userview "projek/be8/delivery/view/user"
	ruser "projek/be8/repository/user"

	"net/http"
	"projek/be8/entities"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type UserController struct {
	Repo  ruser.User
	Valid *validator.Validate
}

func New(repo ruser.User, valid *validator.Validate) *UserController {
	return &UserController{
		Repo:  repo,
		Valid: valid,
	}
}

// func (uc *UserController) GetAllUser(c echo.Context) error {

// 	res, err := uc.Repo.GetAllUser()

// 	if err != nil {
// 		log.Warn("masalah pada server")
// 		return c.JSON(http.StatusInternalServerError, view.InternalServerError())
// 	}
// 	log.Info("berhasil get all data")
// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"code":    http.StatusOK,
// 		"message": "berhasil get all data",
// 		"status":  true,
// 		"data":    res,
// 	})
// }

func (uc *UserController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		param := userview.LoginRequest{}

		if err := c.Bind(&param); err != nil {
			log.Warn("salah input")
			return c.JSON(http.StatusUnsupportedMediaType, view.BindData())
		}

		if err := uc.Valid.Struct(&param); err != nil {
			log.Warn(err.Error())
			return c.JSON(http.StatusNotAcceptable, view.Validate())
		}

		hasil, err := uc.Repo.Login(param.Email, param.Password)

		if err != nil {
			log.Warn(err.Error())
			return c.JSON(http.StatusNotFound, view.NotFound())
		}

		res := userview.LoginResponse{}

		if res.Token == "" {
			token, _ := middlewares.CreateToken(int(hasil.ID))
			res.Token = token
			return c.JSON(http.StatusOK, userview.LoginOK(res, "Berhasil login"))
		}

		return c.JSON(http.StatusOK, userview.LoginOK(res, "Berhasil login"))
	}
}
func (uc *UserController) InsertUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var tmpUser userview.InsertUserRequest

		if err := c.Bind(&tmpUser); err != nil {
			log.Warn("salah input")
			return c.JSON(http.StatusUnsupportedMediaType, view.BindData())
		}

		if err := uc.Valid.Struct(&tmpUser); err != nil {
			log.Warn(err.Error())
			return c.JSON(http.StatusNotAcceptable, view.Validate())
		}

		newUser := entities.User{
			Name:     tmpUser.Name,
			Email:    tmpUser.Email,
			Password: tmpUser.Password,
			Phone:    tmpUser.Phone}
		res, err := uc.Repo.InsertUser(newUser)

		if err != nil {
			log.Warn("masalah pada server")
			return c.JSON(http.StatusInternalServerError, view.InternalServerError())
		}
		log.Info("berhasil insert")
		return c.JSON(http.StatusCreated, userview.SuccessInsert(res))
	}
}
