package entities

import "gorm.io/gorm"

type Produk struct {
	gorm.Model
	UserID       int         `json:"user_id"`
	Nama         string      `json:"nama"`
	Stok         int         `json:"stok"`
	Transaksi_id []Transaksi `json:"produk_id" gorm:"foreignKey:produk_id"`
}
