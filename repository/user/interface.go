package user

import "projek/be8/entities"

type User interface {
	InsertUser(newUser entities.User) (entities.User, error)
	Login(email string, password string) (entities.User, error)
}
