package controller

import (
	"agit-backend/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

//create/register user controller
func (ce *EchoController) CreateUserController(c echo.Context) error {
	//bind input to user model
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	//create user from service
	err := ce.Svc.CreateUserService(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	//return success
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success",
	})
}

//login user controller
func (ce *EchoController) LoginUserController(c echo.Context) error {
	//bind input to userLogin model
	userLogin := model.UserLogin{}
	if err := c.Bind(&userLogin); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	//login user from service and get token
	token, statusCode := ce.Svc.LoginUserService(userLogin.Username, userLogin.Password)

	switch statusCode {
	//if username or password incorrect
	case http.StatusUnauthorized:
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"messages": "username atau password salah",
		})

	//if error create token
	case http.StatusInternalServerError:
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": "internal, error create token",
		})
	}

	//return message and token
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success login as " + userLogin.Username,
		"token":    token,
	})
}
