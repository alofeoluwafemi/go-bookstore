package models

import (
	"github.com/slim12kg/go-bookstore/pkg/config"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name string `gorm:"" json:"name"`
	Author string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	if err := db.AutoMigrate(&Book{}); err != nil {
		log.Fatalln("Init error: ", err)
	}
}

func (b *Book) CreateBook() *Book {
	db.Create(&b)
	return b
}

func GetAllBooks() []Book  {
	var books []Book
	db.Find(&books)
	return books
}

func GetBookById(bookId int64)  (*Book, *gorm.DB){
	var book Book
	db := db.Where("ID=?",bookId).Find(&book)
	return &book,db
}

func DeleteBookById(bookId int64)  (*Book, *gorm.DB){
	var book Book
	db := db.Where("ID=?",bookId).Find(&book).Delete(&book)
	return &book,db
}

func UpdateBookById(bookId int64, fields map[string]interface{}) (*Book, *gorm.DB) {
	var book Book

	db := db.Where("ID=?",bookId).Find(&book).Updates(fields)

	return &book, db
}