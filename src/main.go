package main

import (
	"fmt"

	"github.com/glebarez/sqlite"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID       uint `gorm:"primaryKey"`
	Username string
	Email    string
	Password string
}

func FetchAll(db *gorm.DB, users *[]User) *gorm.DB {
	return db.Find(&users)
}

func CreateUser(db *gorm.DB) *gorm.DB {
	password, err := bcrypt.GenerateFromPassword([]byte("Krneki123"), 9)

	if err != nil {
		fmt.Println(err)
	}

	user := User{Username: "frasza", Email: "zan.fraas@gmail.com", Password: string(password)}

	return db.Create(&user)
}

func main() {
	var users []User
	db, err := gorm.Open(sqlite.Open("./src/database.db"), &gorm.Config{})

	if err != nil {
		fmt.Println("Could not connect to db.")
	}

	// CreateUser(db)

	result := db.Find(&users)

	fmt.Println(result)
}
