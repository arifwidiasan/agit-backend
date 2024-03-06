package service

import (
	"fmt"
	"net/http"

	"agit-backend/helper"
	"agit-backend/model"

	"golang.org/x/crypto/bcrypt"
)

//create user service
func (s *svc) CreateUserService(user model.User) error {
	//check if username already exist
	_, err := s.repo.GetUserByUsername(user.Username)
	if err == nil {
		return fmt.Errorf("username already exist")
	}

	//check if username and password is empty
	if user.Username == "" || user.Password == "" {
		return fmt.Errorf("field username and password cannot be empty")
	}

	//generate password
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("error generate password")
	}
	user.Password = string(hash) //change password to hash

	//insert user to repository
	err = s.repo.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}

//login user service
func (s *svc) LoginUserService(username, password string) (string, int) {
	//get user by username
	user, err := s.repo.GetUserByUsername(username)
	if err != nil {
		return "", http.StatusUnauthorized
	}

	//compare password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", http.StatusUnauthorized
	}

	//generate jwt token
	token, err := helper.CreateToken(user.ID, user.Username, s.c.JWT_KEY)
	if err != nil {
		return "", http.StatusInternalServerError
	}

	return token, http.StatusOK
}
