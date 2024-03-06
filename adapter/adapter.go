package adapter

import (
	"agit-backend/model"

	"github.com/golang-jwt/jwt"
)

type AdapterRepository interface {
	//adapter user repository
	CreateUser(user model.User) error
	GetUserByUsername(username string) (user model.User, err error)
	//adapter karyawan repository
	CreateKaryawan(karyawan model.Karyawan) error
	GetAllKaryawan() []model.Karyawan
	GetKaryawanByID(id int) (karyawan model.Karyawan, err error)
	UpdateKaryawanByID(id int, karyawan model.Karyawan) error
	DeleteKaryawanByID(id int) error
	GetAllSoftDeleteKaryawan() []model.Karyawan
	RestoreSoftDeleteKaryawan(id int) error
}

type AdapterService interface {
	ClaimToken(bearer *jwt.Token) string
	//adapter user service
	LoginUserService(username, password string) (string, int)
	CreateUserService(user model.User) error
	GetUserByUsernameService(username string) (user model.User, err error)
	//adapter karyawan service
	CreateKaryawanService(karyawan model.Karyawan) error
	GetAllKaryawanService() []model.Karyawan
	GetKaryawanByIDService(id int) (karyawan model.Karyawan, err error)
	UpdateKaryawanByIDService(id int, karyawan model.Karyawan) error
	DeleteKaryawanByIDService(id int) error
	GetAllSoftDeleteKaryawanService() []model.Karyawan
	RestoreSoftDeleteKaryawanService(id int) error
}
