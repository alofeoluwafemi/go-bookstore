package routes

import "github.com/gorilla/mux"
import "github.com/slim12kg/go-bookstore/pkg/controllers"

var RegisterBookstoreRoutes = func(route *mux.Router) {
	route.HandleFunc("/book",controllers.CreateBook).Methods("POST")
	route.HandleFunc("/books",controllers.GetBooks).Methods("GET")
	route.HandleFunc("/book/{bookId}",controllers.GetBookById).Methods("GET")
	route.HandleFunc("/book/{bookId}",controllers.UpdateBook).Methods("PUT")
	route.HandleFunc("/book/{bookId}",controllers.DeleteBook).Methods("DELETE")
}