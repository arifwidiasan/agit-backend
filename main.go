package main

import (
	conf "agit-backend/config"
	handler "agit-backend/controller/handler"

	"github.com/labstack/echo/v4"
)

func main() {
	config := conf.InitConfiguration()
	e := echo.New()

	handler.RegisterGroupAPI(e, config)

	e.Logger.Fatal(e.Start((config.SERVER_ADDRESS)))
}
