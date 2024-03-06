package adapter

import (
	"agit-backend/model"
)

type AdapterRepository interface {
	CreateUser(user model.User) error
	GetUserByUsername(username string) (user model.User, err error)
}

type AdapterService interface {
	CreateUserService(user model.User) error
	LoginUserService(username, password string) (string, int)
}
