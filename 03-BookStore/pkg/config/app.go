package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	dsn := "admin:admin123@tcp(127.0.0.1:3306)/books?charset=utf8mb4&parseTime=True&loc=Local"

	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {

		fmt.Println("Cannot connect to mysql db")
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
