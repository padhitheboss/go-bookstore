package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	//"github.com/jinzhu/gorm/dialects/mysql"
	"github.com/padhitheboss/go-bookstore/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	fmt.Println("Starting Server on Port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
