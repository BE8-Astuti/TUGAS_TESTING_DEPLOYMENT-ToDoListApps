package transaksi

import "projek/be8/entities"

type Transaksi interface {
	Insert(newTransaksi entities.Transaksi) (entities.Transaksi, error)
	GetAll() ([]entities.Transaksi, error)
	GetTrans(jenis_transaksi string) ([]entities.Transaksi, error)
	HistoriTrans(jenis_transaksi string) ([]entities.Transaksi, error)
	RiwayatAllTrans() ([]entities.Transaksi, error)
}
