package produk

import "projek/be8/entities"

type ProdukRepository interface {
	InsertProduk(newProduk entities.Produk) (entities.Produk, error)
	GetAllProduk() ([]entities.Produk, error)
	GetProdukID(ID int) (entities.Produk, error)
	UpdateProduk(id int, UpdateProduk entities.Produk, UserID int) (entities.Produk, error)
	DeleteProduk(id int, UserID int) error
}
