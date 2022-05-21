package user

import (
	"errors"
	"projek/be8/entities"

	"github.com/labstack/gommon/log"

	"gorm.io/gorm"
)

type UserRepo struct {
	Db *gorm.DB
}

func New(db *gorm.DB) *UserRepo {
	return &UserRepo{
		Db: db,
	}
}

func (ur *UserRepo) InsertUser(newUser entities.User) (entities.User, error) {
	if err := ur.Db.Create(&newUser).Error; err != nil {
		log.Warn(err)
		return entities.User{}, errors.New("tidak bisa insert data")
	}

	log.Info()
	return newUser, nil
}

func (ur *UserRepo) Login(email string, password string) (entities.User, error) {
	users := []entities.User{}

	if err := ur.Db.Where("email = ? AND password = ?", email, password).First(&users).Error; err != nil {
		log.Warn(err)
		return entities.User{}, errors.New("tidak bisa select data")
	}

	return users[0], nil
}
