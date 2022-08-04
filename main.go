package main

import (
	"foodmarketplace/config"
	"foodmarketplace/database"
	"foodmarketplace/routes"
	"github.com/labstack/echo/v4"
)

func main() {

	config.SetConfig()
	database.Connect()

	e := echo.New()
	routes.SetRoutes(e)
	e.Start(config.AppConfig.Port)

}
