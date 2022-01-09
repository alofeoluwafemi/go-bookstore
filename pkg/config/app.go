package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var
(
	db *gorm.DB
)

func Connect() {
	dsn := "root:Swagger@tcp(127.0.0.1:3306)/bookstore?charset=utf8&parseTime=True&loc=Local"
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("Connection error: ", err)
	}
	
	db = conn
}

func GetDB() *gorm.DB {
	return db
}