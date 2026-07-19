package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db * gorm.DB
)

func Connect(){
	// first download mysql server and install it, then create a database named bookstore, then run the following command to create a user and add user and pass here in the connection string below
	d, err := gorm.Open("mysql", "username:password@tcp(localhost:3306)/bookstore?charset=utf8mb4&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}

	db = d
}

func GetDB() *gorm.DB{
	return db
}