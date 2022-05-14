package entities

import "gorm.io/gorm"

type Transaksi struct {
	gorm.Model
	Produk_id       int    `json:"produk_id"`
	User_id         int    `json:"user_id"`
	Produk          string `json:"produk"`
	Qty             int    `json:"qty"`
	Jenis_transaksi string `json:"jenis_transaksi"`
}
