package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	middlewares "projek/be8/delivery/middlewares"
	"projek/be8/entities"
	"strings"
	"testing"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/stretchr/testify/assert"
)

var token string

func TestCreateToken(t *testing.T) {
	t.Run("Create Token", func(t *testing.T) {
		token, _ = middlewares.CreateToken(1)
	})
}

func TestInsertUser(t *testing.T) {
	t.Run("Success Insert", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":     "yani",
			"email":    "y",
			"password": "849",
			"phone":    "77979799",
		})
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON) // Set Content to JSON

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/user")

		userController := New(&mockUserRepository{}, validator.New())
		userController.InsertUser()(context)

		type response struct {
			Code    int
			Message string
			Status  bool
			Data    interface{}
		}

		var resp response

		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		assert.Equal(t, "berhasil insert data user", resp.Message)
		assert.True(t, resp.Status)
		assert.Equal(t, 201, resp.Code)
		assert.NotNil(t, resp.Data)
	})
	t.Run("Error Validasi", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":     "yani",
			"password": "849",
			"phone":    "77979799",
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/user")

		userController := New(&erorrMockUserRepository{}, validator.New())
		userController.InsertUser()(context)

		type response struct {
			Code    int
			Message string
			Status  bool
			Data    interface{}
		}

		var resp response

		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		log.Warn(resp)
		assert.False(t, resp.Status)
		assert.Nil(t, resp.Data)
		assert.Equal(t, 406, resp.Code)
	})
	t.Run("Error Bind", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"phone": "779",
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/user")

		userController := New(&erorrMockUserRepository{}, validator.New())
		userController.InsertUser()(context)

		type response struct {
			Code    int
			Message string
			Status  bool
			Data    interface{}
		}

		var resp response

		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		log.Warn(resp)
		assert.False(t, resp.Status)
		assert.Nil(t, resp.Data)
		assert.Equal(t, 415, resp.Code)
	})
	t.Run("Error Insert DB", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":     "yani",
			"email":    "y",
			"password": "849",
			"phone":    "77979799",
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/user")

		userController := New(&erorrMockUserRepository{}, validator.New())
		userController.InsertUser()(context)

		type response struct {
			Code    int
			Message string
			Status  bool
			Data    interface{}
		}

		var resp response

		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		assert.False(t, resp.Status)
		assert.Nil(t, resp.Data)
		assert.Equal(t, 500, resp.Code)
	})
}

func TestLogin(t *testing.T) {
	t.Run("Success Login", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"email":    "y@gmail.com",
			"password": "yani",
		})
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/login")

		controller := New(&mockUserRepository{}, validator.New())
		controller.Login()(context)

		type ResponseStructure struct {
			Code    int
			Message string
			Status  bool
			Data    interface{}
		}

		var response ResponseStructure

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, 200, response.Code)
		assert.True(t, response.Status)
		assert.NotNil(t, response.Data)
		data := response.Data.(map[string]interface{})
		token = data["Token"].(string)
	})
	t.Run("Error Validasi", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"password": "779",
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/login")

		userController := New(&erorrMockUserRepository{}, validator.New())
		userController.Login()(context)

		type response struct {
			Code    int
			Message string
			Status  bool
			Data    interface{}
		}

		var resp response

		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		log.Warn(resp)
		assert.False(t, resp.Status)
		assert.Nil(t, resp.Data)
		assert.Equal(t, 406, resp.Code)
	})
	t.Run("Error Bind", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"password": "779",
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/login")

		userController := New(&erorrMockUserRepository{}, validator.New())
		userController.Login()(context)

		type response struct {
			Code    int
			Message string
			Status  bool
			Data    interface{}
		}

		var resp response

		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		log.Warn(resp)
		assert.False(t, resp.Status)
		assert.Nil(t, resp.Data)
		assert.Equal(t, 415, resp.Code)
	})
	t.Run("Error Login", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"email":    "y@gmail.com",
			"password": "yani",
		})
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/login")

		controller := New(&erorrMockUserRepository{}, validator.New())
		controller.Login()(context)

		type Response struct {
			Code    int
			Message string
			Status  bool
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 404, result.Code)
		assert.Equal(t, "Data Not Found", result.Message)
		assert.False(t, result.Status)
	})
}

type mockUserRepository struct{}

func (mur *mockUserRepository) InsertUser(newUser entities.User) (entities.User, error) {
	return entities.User{Name: "Astuti", Email: "a@gmail.com", Phone: "7897787", Jenis_transaksi: "pembelian"}, nil
}
func (mur *mockUserRepository) Login(email, password string) (entities.User, error) {
	return entities.User{Name: "Astuti", Email: "a@gmail.com", Phone: "7897787", Jenis_transaksi: "pembelian"}, nil
}

type erorrMockUserRepository struct{}

func (emur *erorrMockUserRepository) InsertUser(newPegawai entities.User) (entities.User, error) {
	return entities.User{}, errors.New("tidak bisa insert data")
}

func (emur *erorrMockUserRepository) Login(email, password string) (entities.User, error) {
	return entities.User{}, errors.New("tidak bisa select data")
}
