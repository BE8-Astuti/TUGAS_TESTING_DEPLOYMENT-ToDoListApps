package produk

import (
	"net/http"
	"projek/be8/entities"
)

func StatusGetAllOk(data []entities.Produk) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success Get All data",
		"status":  true,
		"data":    data,
	}
}

func StatusGetIdOk(data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success Get Data",
		"status":  true,
		"data":    data,
	}
}

func StatusCreate(data entities.Produk) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusCreated,
		"message": "Success Create Product",
		"status":  true,
		"data":    data,
	}
}

func StatusUpdate(data entities.Produk) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Updated",
		"status":  true,
		"data":    data,
	}
}
