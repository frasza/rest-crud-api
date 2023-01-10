package router

import (
	"crud-rest-api/src/handlers"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	app := echo.New()

	app.GET("/users", handlers.GetUsers)
	app.POST("/users", handlers.CreateUser)

	return app
}
