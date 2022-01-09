package main

import (
	"github.com/gorilla/mux"
	"github.com/slim12kg/go-bookstore/pkg/routes"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookstoreRoutes(r)
	http.Handle("/",r)
	if err := http.ListenAndServe("localhost:8080",r); err != nil{
		log.Fatalln("Error running server: ", err)
	}
}