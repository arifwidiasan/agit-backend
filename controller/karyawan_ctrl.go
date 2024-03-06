package controller

import (
	"agit-backend/model"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

//create karyawan controller
func (c *EchoController) CreateKaryawanController(e echo.Context) error {
	//check JWT token
	user := c.Svc.ClaimToken(e.Get("user").(*jwt.Token))
	_, err := c.Svc.GetUserByUsernameService(user)
	if err != nil {
		return e.JSON(http.StatusUnauthorized, map[string]string{
			"message": err.Error(),
		})
	}

	//bind request body to model karyawan
	karyawan := model.Karyawan{}
	err = e.Bind(&karyawan)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	//create karyawan
	err = c.Svc.CreateKaryawanService(karyawan)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]string{
		"message": "success",
	})
}

//get all karyawan controller
func (c *EchoController) GetAllKaryawanController(e echo.Context) error {
	//check JWT token
	user := c.Svc.ClaimToken(e.Get("user").(*jwt.Token))
	_, err := c.Svc.GetUserByUsernameService(user)
	if err != nil {
		return e.JSON(http.StatusUnauthorized, map[string]string{
			"message": err.Error(),
		})
	}

	//get all karyawan
	karyawans := c.Svc.GetAllKaryawanService()

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    karyawans,
	})
}

//get karyawan by id controller
func (c *EchoController) GetKaryawanByIDController(e echo.Context) error {
	//check JWT token
	user := c.Svc.ClaimToken(e.Get("user").(*jwt.Token))
	_, err := c.Svc.GetUserByUsernameService(user)
	if err != nil {
		return e.JSON(http.StatusUnauthorized, map[string]string{
			"message": err.Error(),
		})
	}

	//get karyawan by id
	id, _ := strconv.Atoi(e.Param("id"))
	karyawan, err := c.Svc.GetKaryawanByIDService(id)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    karyawan,
	})
}

//update karyawan by id controller
func (c *EchoController) UpdateKaryawanByIDController(e echo.Context) error {
	//check JWT token
	user := c.Svc.ClaimToken(e.Get("user").(*jwt.Token))
	_, err := c.Svc.GetUserByUsernameService(user)
	if err != nil {
		return e.JSON(http.StatusUnauthorized, map[string]string{
			"message": err.Error(),
		})
	}

	//bind request body to model karyawan
	karyawan := model.Karyawan{}
	err = e.Bind(&karyawan)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	//update karyawan by id
	id, _ := strconv.Atoi(e.Param("id"))
	err = c.Svc.UpdateKaryawanByIDService(id, karyawan)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]string{
		"message": "success",
	})
}
