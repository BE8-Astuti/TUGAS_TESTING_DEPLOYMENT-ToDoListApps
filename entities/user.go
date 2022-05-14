package entities

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name            string      `json:"name"`
	Password        string      `gorm:"unique" json:"password" form:"password"`
	Email           string      `json:"email"`
	Phone           string      `json:"phone"`
	Jenis_transaksi string      `json:"jenis_transaksi"`
	Transaksi_id    []Transaksi `gorm:"foreignKey:user_id"`
}
