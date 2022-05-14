package produk

import (
	"errors"
	"projek/be8/entities"

	"github.com/labstack/gommon/log"

	"gorm.io/gorm"
)

type ProdukRepo struct {
	Db *gorm.DB
}

func New(db *gorm.DB) *ProdukRepo {
	return &ProdukRepo{
		Db: db,
	}
}

func (pr *ProdukRepo) InsertProduk(newProduk entities.Produk) (entities.Produk, error) {
	if err := pr.Db.Create(&newProduk).Error; err != nil {
		log.Warn(err)
		return entities.Produk{}, errors.New("tidak bisa insert produk")
	}

	log.Info()
	return newProduk, nil
}

func (pr *ProdukRepo) GetAllProduk() ([]entities.Produk, error) {
	arrProduk := []entities.Produk{}

	if err := pr.Db.Find(&arrProduk).Error; err != nil {
		log.Warn(err)
		return nil, errors.New("tidak bisa select data produk")
	}

	if len(arrProduk) == 0 {
		log.Warn("tidak ada data")
		return nil, errors.New("tidak ada data")
	}

	log.Info()
	return arrProduk, nil
}
func (pr *ProdukRepo) GetProdukID(ID int) (entities.Produk, error) {
	arrProduk := []entities.Produk{}

	if err := pr.Db.Where("id = ?", ID).Find(&arrProduk).Error; err != nil {
		log.Warn(err)
		return entities.Produk{}, errors.New("tidak bisa select data")
	}

	if len(arrProduk) == 0 {
		log.Warn("data tidak ditemukan")
		return entities.Produk{}, errors.New("data tidak ditemukan")
	}

	log.Info()
	return arrProduk[0], nil
}

func (pr *ProdukRepo) UpdateProduk(id int, UpdateProduk entities.Produk, UserID int) (entities.Produk, error) {
	var produks entities.Produk
	if err := pr.Db.Where("id =? AND user_id =?", id, UserID).First(&produks).Updates(&UpdateProduk).Find(&produks).Error; err != nil {
		log.Warn("Update Address Error", err)
		return produks, err
	}

	return produks, nil
}

func (pr *ProdukRepo) DeleteProduk(id int, UserID int) error {

	var delete entities.Produk
	if err := pr.Db.Where("id = ? AND user_id = ?", id, UserID).First(&delete).Delete(&delete).Error; err != nil {
		return err
	}
	return nil
}
