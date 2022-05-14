package transaksi

type InsertTransaksiRequest struct {
	User_id         int    `json:"user_id" validate:"required"`
	Produk_id       int    `json:"produk_id" validate:"required"`
	Produk          string `json:"produk" validate:"required"`
	Qty             int    `json:"qty"`
	Jenis_transaksi string `json:"jenis_transaksi"`
}
