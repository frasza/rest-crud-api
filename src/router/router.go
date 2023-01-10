package router

import (
	"crud-rest-api/src/handlers"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	app := echo.New()

	app.GET("/users", handlers.GetUsers)
	app.GET("/users/:id", handlers.GetSingleUser)
	app.POST("/users", handlers.CreateUser)
	app.PUT("/users/:id", handlers.UpdateUser)
	app.DELETE("/users/:id", handlers.DeleteUser)

	return app
}
