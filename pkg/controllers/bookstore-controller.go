package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/slim12kg/go-bookstore/pkg/models"
	"github.com/slim12kg/go-bookstore/pkg/utils"
	"log"
	"net/http"
	"strconv"
)

var books [] models.Book

func GetBooks(w http.ResponseWriter, req *http.Request) {
	books = models.GetAllBooks()
	respondData(books, w)
}

func GetBookById(w http.ResponseWriter, req *http.Request)  {
	Id := getBookId(req)
	book , _ := models.GetBookById(Id)

	respondData(book, w)
}

func CreateBook(w http.ResponseWriter, req *http.Request)  {
	bookModel := &models.Book{}
	utils.ParseBody(req, bookModel)
	book := bookModel.CreateBook()

	respondData(book, w)
}

func UpdateBook(w http.ResponseWriter, req *http.Request) {
	bookModel := &models.Book{}
	utils.ParseBody(req, bookModel)
	Id := getBookId(req)

	updatedBook := map[string]interface{}{"Name": bookModel.Name,"Author": bookModel.Author, "Publication": bookModel.Publication}

	book , _ := models.UpdateBookById(Id,updatedBook)

	respondData(book, w)
}

func DeleteBook(w http.ResponseWriter, req *http.Request) {
	Id := getBookId(req)
	book , _ := models.DeleteBookById(Id)

	respondData(book, w)
}

func getBookId(req *http.Request) int64 {
	vars := mux.Vars(req)
	bookId := vars["bookId"]
	Id, err := strconv.ParseInt(bookId,10,64)

	if err != nil {
		log.Fatalln(err)
	}

	return Id
}
func respondData(body interface{}, w http.ResponseWriter){
	data, err := json.Marshal(body)
	if err != nil {
		log.Fatalln(err)
	}

	w.Header().Set("Content-Type","application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(data); err != nil {
		log.Fatalln("Error writing error: ", err)
	}
}