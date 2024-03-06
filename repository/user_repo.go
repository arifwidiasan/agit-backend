package repository

import (
	"fmt"

	"agit-backend/model"
)

//create user
func (r *repositoryMysqlLayer) CreateUser(user model.User) error {
	//insert user
	res := r.DB.Create(&user)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error create user or username already exist")
	}
	return nil
}

//get user by username
func (r *repositoryMysqlLayer) GetUserByUsername(username string) (user model.User, err error) {
	res := r.DB.Where("username = ?", username).Find(&user)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("user not found")
	}

	return
}
