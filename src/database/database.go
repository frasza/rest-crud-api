package database

import (
	"crud-rest-api/src/model"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func Init() {
	db, err = gorm.Open(sqlite.Open("./src/database/database.db"), &gorm.Config{})
	if err != nil {
		panic("Could not connect to DB. Panic time!")
	}

	db.AutoMigrate(&model.User{})
}

func DatabaseManager() *gorm.DB {
	return db
}
