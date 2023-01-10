package handlers

import (
	"crud-rest-api/src/database"
	"crud-rest-api/src/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func GetUsers(c echo.Context) error {
	db := database.DatabaseManager()

	users := []model.User{}
	db.Find(&users)

	return c.JSON(http.StatusOK, users)
}

func GetSingleUser(c echo.Context) error {
	db := database.DatabaseManager()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	user := &model.User{}

	db.First(&user, id)

	return c.JSON(http.StatusOK, user)
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

	return c.JSON(http.StatusCreated, user)
}

func UpdateUser(c echo.Context) error {
	db := database.DatabaseManager()
	user := &model.User{}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	db.First(&user, id)

	updated := &model.User{}

	if err := c.Bind(updated); err != nil {
		return err
	}

	user.Username = updated.Username
	user.Email = updated.Email

	db.Save(&user)

	return c.JSON(http.StatusOK, user)
}

func DeleteUser(c echo.Context) error {
	db := database.DatabaseManager()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	user := &model.User{}

	db.First(&user, id)

	db.Delete(&user)

	return c.NoContent(http.StatusNoContent)
}
