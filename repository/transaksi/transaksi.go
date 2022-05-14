package transaksi

import (
	"errors"
	"projek/be8/entities"

	"github.com/labstack/gommon/log"

	"gorm.io/gorm"
)

type TransaksiRepo struct {
	Db *gorm.DB
}

func New(db *gorm.DB) *TransaksiRepo {
	return &TransaksiRepo{
		Db: db,
	}
}

func (tr *TransaksiRepo) Insert(newTransaksi entities.Transaksi) (entities.Transaksi, error) {
	if err := tr.Db.Create(&newTransaksi).Error; err != nil {
		log.Warn(err)
		return entities.Transaksi{}, errors.New("tidak bisa insert data")
	}

	log.Info()
	return newTransaksi, nil
}

func (tr *TransaksiRepo) GetAll() ([]entities.Transaksi, error) {
	arrTransaksi := []entities.Transaksi{}

	if err := tr.Db.Find(&arrTransaksi).Error; err != nil {
		log.Warn(err)
		return nil, errors.New("tidak bisa select data")
	}

	if len(arrTransaksi) == 0 {
		log.Warn("tidak ada data")
		return nil, errors.New("tidak ada data")
	}

	log.Info()
	return arrTransaksi, nil
}

func (tr *TransaksiRepo) HistoriTrans(jenis_transaksi string) ([]entities.Transaksi, error) {
	arrTransaksi := []entities.Transaksi{}

	if err := tr.Db.Where("jenis_transaksi = ? AND created_at BETWEEN CURDATE()-7 AND CURDATE()", jenis_transaksi).Find(&arrTransaksi).Error; err != nil {
		log.Warn(err)
		return nil, errors.New("tidak bisa select data")
	}

	if len(arrTransaksi) == 0 {
		log.Warn("tidak ada data")
		return nil, errors.New("tidak ada data")
	}

	log.Info()
	return arrTransaksi, nil
}

func (tr *TransaksiRepo) RiwayatAllTrans() ([]entities.Transaksi, error) {
	arrTransaksi := []entities.Transaksi{}

	if err := tr.Db.Where("created_at BETWEEN CURDATE()-7 AND CURDATE()").Find(&arrTransaksi).Error; err != nil {
		log.Warn(err)
		return nil, errors.New("tidak bisa select data")
	}

	if len(arrTransaksi) == 0 {
		log.Warn("tidak ada data")
		return nil, errors.New("tidak ada data")
	}

	log.Info()
	return arrTransaksi, nil
}

func (tr *TransaksiRepo) GetTrans(jenis_transaksi string) ([]entities.Transaksi, error) {
	arrTransaksi := []entities.Transaksi{}

	if err := tr.Db.Where("jenis_transaksi = ?", jenis_transaksi).Find(&arrTransaksi).Error; err != nil {
		log.Warn(err)
		return arrTransaksi, errors.New("tidak bisa select data")
	}
	if len(arrTransaksi) == 0 {
		log.Warn("data tidak ditemukan")
		return arrTransaksi, errors.New("data tidak ditemukan")
	}

	log.Info()
	return arrTransaksi, nil

}
