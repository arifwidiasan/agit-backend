package handler

import (
	"agit-backend/config"
	"agit-backend/controller"
	"agit-backend/database"

	m "agit-backend/middleware"
	"agit-backend/repository"
	"agit-backend/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterGroupAPI(e *echo.Echo, conf config.Config) {

	db := database.InitDB(conf)
	repo := repository.NewMysqlRepository(db)

	svc := service.NewService(repo, conf)

	ctrl := controller.EchoController{
		Svc: svc,
	}

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(200, map[string]string{
			"message": "your request awesome",
		})
	})

	api := e.Group("/api", middleware.CORS())

	m.LogMiddleware(e)
	api.POST("/register", ctrl.CreateUserController)

	api.POST("/login", ctrl.LoginUserController)
}
