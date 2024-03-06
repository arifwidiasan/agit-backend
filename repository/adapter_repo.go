package repository

import (
	"agit-backend/adapter"

	"gorm.io/gorm"
)

type repositoryMysqlLayer struct {
	DB *gorm.DB
}

func NewMysqlRepository(db *gorm.DB) adapter.AdapterRepository {
	return &repositoryMysqlLayer{
		DB: db,
	}
}
