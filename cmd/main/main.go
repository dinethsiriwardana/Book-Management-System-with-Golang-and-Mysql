package main

import (
	"log"
	"net/http"

	"github.com/dinethsiriwardana/Book-Management-System-with-Golang-and-Mysql/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRouters(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9090", r))
}
