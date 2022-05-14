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
	"github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
)

var token string

// INITIATE TOKEN
func TestCreateToken(t *testing.T) {
	t.Run("Create Token", func(t *testing.T) {
		token, _ = middlewares.CreateToken(1)
	})
}

func TestInsert(t *testing.T) {
	t.Run("Create Success", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"user_id":         1,
			"produk_id":       1,
			"produk":          "tango",
			"qty":             10,
			"jenis_transaksi": "penjualan",
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/produk")
		InserTran := New(&mockTransaksi{}, &mockProduct{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("A$T0T!")})(InserTran.InsertTransaksi())(context)

		type Response struct {
			Code    int
			Message string
			Status  bool
			Data    interface{}
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 201, result.Code)
		assert.Equal(t, "Success Create Transaksi", result.Message)
		assert.True(t, result.Status)
		assert.NotNil(t, result.Data)
	})
	t.Run("Error Access Database", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"user_id":         1,
			"produk_id":       1,
			"produk":          "tango",
			"qty":             10,
			"jenis_transaksi": "penjualan",
		})
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/produk")
		transaksi := New(&errMockTransaksi{}, &errMockProduct{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("A$T0T!")})(transaksi.InsertTransaksi())(context)

		type Response struct {
			Code    int
			Message string
			Status  bool
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 500, result.Code)
		assert.Equal(t, "Cannot Access Database", result.Message)
		assert.False(t, result.Status)
	})
	t.Run("Error Bind", func(t *testing.T) {
		e := echo.New()

		requestBody := "kecantikan"

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/transaksi")
		Prod := New(&errMockTransaksi{}, &errMockProduct{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("A$T0T!")})(Prod.InsertTransaksi())(context)

		type Response struct {
			Code    int
			Message string
			Status  bool
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)
		assert.Equal(t, 415, result.Code)
		assert.Equal(t, "Cannot Bind Data", result.Message)
		assert.False(t, result.Status)
	})
	t.Run("Error Validate", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{

			"produk_id":       1,
			"produk":          "tango",
			"qty":             10,
			"jenis_transaksi": "penjualan",
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/transaksi")
		Prod := New(&errMockTransaksi{}, &errMockProduct{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("A$T0T!")})(Prod.InsertTransaksi())(context)

		type Response struct {
			Code    int
			Message string
			Status  bool
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 406, result.Code)
		assert.Equal(t, "Validate Error", result.Message)
		assert.False(t, result.Status)
	})

}

func TestGetAllTransaksi(t *testing.T) {
	t.Run("Success Get All Transaksi", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/transaksi")
		GetProd := New(&mockTransaksi{}, &mockProduct{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("A$T0T!")})(GetProd.GetAllTransaksi())(context)

		type Response struct {
			Code    int
			Message string
			Status  bool
			Data    interface{}
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 200, result.Code)
		assert.Equal(t, "Success Get All data", result.Message)
		assert.True(t, result.Status)
		assert.NotNil(t, result.Data)
	})
	t.Run("Error Access Database", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/transaksi")
		GetTra := New(&errMockTransaksi{}, &errMockProduct{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("A$T0T!")})(GetTra.GetAllTransaksi())(context)

		type Response struct {
			Code    int
			Message string
			Status  bool
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 500, result.Code)
		assert.Equal(t, "Cannot Access Database", result.Message)
		assert.False(t, result.Status)
	})
}

func TestGetTransaksi(t *testing.T) {
	t.Run("Success Get All Transaksi", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/transaksi/:tipe")
		context.SetParamNames("tipe")
		context.SetParamValues("pembelian")
		GetProd := New(&mockTransaksi{}, &mockProduct{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("A$T0T!")})(GetProd.GetTransaksi())(context)

		type Response struct {
			Code    int
			Message string
			Status  bool
			Data    interface{}
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 200, result.Code)
		assert.Equal(t, "Success Get All data", result.Message)
		assert.True(t, result.Status)
		assert.NotNil(t, result.Data)
	})
	t.Run("Error Access Database", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/transaksi/:tipe")
		context.SetParamNames("tipe")
		context.SetParamValues("pembelian")
		GetTra := New(&errMockTransaksi{}, &errMockProduct{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("A$T0T!")})(GetTra.RiwayatAllTrans())(context)

		type Response struct {
			Code    int
			Message string
			Status  bool
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 500, result.Code)
		assert.Equal(t, "Cannot Access Database", result.Message)
		assert.False(t, result.Status)
	})
}

func TestRiwayatAllTrans(t *testing.T) {
	t.Run("Success Get All Transaksi", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/transaksi/recently7days")
		GetProd := New(&mockTransaksi{}, &mockProduct{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("A$T0T!")})(GetProd.RiwayatAllTrans())(context)

		type Response struct {
			Code    int
			Message string
			Status  bool
			Data    interface{}
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 200, result.Code)
		assert.Equal(t, "Success Get All data", result.Message)
		assert.True(t, result.Status)
		assert.NotNil(t, result.Data)
	})
	t.Run("Error Access Database", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/transaksi/recently7days")
		GetTra := New(&errMockTransaksi{}, &errMockProduct{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("A$T0T!")})(GetTra.RiwayatAllTrans())(context)

		type Response struct {
			Code    int
			Message string
			Status  bool
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 500, result.Code)
		assert.Equal(t, "Cannot Access Database", result.Message)
		assert.False(t, result.Status)
	})
}
func TestHistoriTrans(t *testing.T) {

	t.Run("Success Get All Transaksi", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/transaksi/:tipe/last-7-days")
		context.SetParamNames("tipe")
		context.SetParamValues("pembelian")
		GetProd := New(&mockTransaksi{}, &mockProduct{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("A$T0T!")})(GetProd.HistoriTrans())(context)

		type Response struct {
			Code    int
			Message string
			Status  bool
			Data    interface{}
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 200, result.Code)
		assert.Equal(t, "Success Get All data", result.Message)
		assert.True(t, result.Status)
		assert.NotNil(t, result.Data)
	})
	t.Run("Error Access Database", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)

		context.SetPath("/transaksi/:tipe/last-7-days")
		context.SetParamNames("tipe")
		context.SetParamValues("pembelian")
		GetTra := New(&errMockTransaksi{}, &errMockProduct{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("A$T0T!")})(GetTra.RiwayatAllTrans())(context)

		type Response struct {
			Code    int
			Message string
			Status  bool
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 500, result.Code)
		assert.Equal(t, "Cannot Access Database", result.Message)
		assert.False(t, result.Status)
	})

}

// MOCK SUCCESS
type mockTransaksi struct {
}

func (s *mockTransaksi) Insert(newTransaksi entities.Transaksi) (entities.Transaksi, error) {
	return entities.Transaksi{Produk_id: 1, User_id: 1, Produk: "tango", Qty: 10, Jenis_transaksi: "penjualan"}, nil
}
func (s *mockTransaksi) GetAll() ([]entities.Transaksi, error) {
	return []entities.Transaksi{{Produk_id: 1, User_id: 1, Produk: "tango", Qty: 10, Jenis_transaksi: "penjualan"}}, nil
}
func (s *mockTransaksi) GetTrans(jenis_transaksi string) ([]entities.Transaksi, error) {
	return []entities.Transaksi{{Produk_id: 1, User_id: 1, Produk: "tango", Qty: 10, Jenis_transaksi: "penjualan"}}, nil
}
func (s *mockTransaksi) HistoriTrans(jenis_transaksi string) ([]entities.Transaksi, error) {
	return []entities.Transaksi{{Produk_id: 1, User_id: 1, Produk: "tango", Qty: 10, Jenis_transaksi: "penjualan"}}, nil
}
func (s *mockTransaksi) RiwayatAllTrans() ([]entities.Transaksi, error) {
	return []entities.Transaksi{{Produk_id: 1, User_id: 1, Produk: "tango", Qty: 10, Jenis_transaksi: "penjualan"}}, nil
}

type mockProduct struct {
}

func (s *mockProduct) InsertProduk(newProduk entities.Produk) (entities.Produk, error) {
	return entities.Produk{Nama: "NIVEA", Stok: 10}, nil
}
func (s *mockProduct) GetAllProduk() ([]entities.Produk, error) {
	return []entities.Produk{{Nama: "NIVEA", Stok: 10}}, nil
}
func (s *mockProduct) GetProdukID(ID int) (entities.Produk, error) {
	return entities.Produk{Nama: "NIVEA", Stok: 10}, nil
}
func (s *mockProduct) UpdateProduk(id int, UpdateProduk entities.Produk, UserID int) (entities.Produk, error) {
	return entities.Produk{Nama: "NIVEA", Stok: 10}, nil
}
func (s *mockProduct) DeleteProduk(id int, UserID int) error {
	return nil
}

type errMockTransaksi struct{}

func (e *errMockTransaksi) Insert(newTransaksi entities.Transaksi) (entities.Transaksi, error) {
	return entities.Transaksi{}, errors.New("Access Database Error")
}
func (e *errMockTransaksi) GetAll() ([]entities.Transaksi, error) {
	return []entities.Transaksi{}, errors.New("Access Database Error")
}
func (e *errMockTransaksi) GetTrans(jenis_transaksi string) ([]entities.Transaksi, error) {
	return []entities.Transaksi{}, errors.New("Access Database Error")
}
func (e *errMockTransaksi) HistoriTrans(jenis_transaksi string) ([]entities.Transaksi, error) {
	return []entities.Transaksi{}, errors.New("Access Database Error")
}
func (e *errMockTransaksi) RiwayatAllTrans() ([]entities.Transaksi, error) {
	return []entities.Transaksi{}, errors.New("Access Database Error")
}

type errMockProduct struct{}

func (d *errMockProduct) InsertProduk(newProduk entities.Produk) (entities.Produk, error) {
	return entities.Produk{}, errors.New("Access Database Error")
}
func (d *errMockProduct) GetAllProduk() ([]entities.Produk, error) {
	return []entities.Produk{}, errors.New("Access Database Error")
}
func (d *errMockProduct) GetProdukID(ID int) (entities.Produk, error) {
	return entities.Produk{}, errors.New("Access Database Error")
}
func (d *errMockProduct) UpdateProduk(id int, UpdateProduk entities.Produk, UserID int) (entities.Produk, error) {
	return entities.Produk{}, errors.New("Access Database Error")
}
func (d *errMockProduct) DeleteProduk(id int, UserID int) error {
	return nil
}
