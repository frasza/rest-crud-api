package handlers

import (
	"crud-rest-api/src/database"
	"crud-rest-api/src/model"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func GetUsers(c echo.Context) error {
	db := database.DatabaseManager()

	users := []model.User{}
	db.Find(&users)

	return c.JSON(http.StatusOK, users)
}

func CreateUser(c echo.Context) error {
	db := database.DatabaseManager()
	user := &model.User{}

	if err := c.Bind(user); err != nil {
		return err
	}

	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return err
	}

	user.Password = string(password)

	db.Create(&user)

	return c.JSON(http.StatusOK, user)
}
