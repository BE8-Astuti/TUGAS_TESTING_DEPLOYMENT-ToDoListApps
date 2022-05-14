package produk

type InsertProdukRequest struct {
	UserID int    `json:"user_id" validate:"required"`
	Nama   string `json:"nama" validate:"required"`
	Stok   int    `json:"stok"`
}

type UpdateProdukRequest struct {
	ID   int    `json:"id" validate:"required"`
	Nama string `json:"nama" validate:"required"`
	Stok int    `json:"stok"`
}
