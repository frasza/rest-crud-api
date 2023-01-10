package main

import (
	"crud-rest-api/src/database"
	"crud-rest-api/src/router"

	"github.com/labstack/echo/v4/middleware"
)

func main() {
	database.Init()

	app := router.Init()

	app.Use(middleware.Logger())

	app.Logger.Fatal(app.Start(":8000"))
}
