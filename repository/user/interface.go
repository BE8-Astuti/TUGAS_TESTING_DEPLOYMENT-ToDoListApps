package user

import "projek/be8/entities"

type User interface {
	InsertUser(newUser entities.User) (entities.User, error)
	Login(name string, password string) (entities.User, error)
}
