package routes

import (
	"github.com/dinethsiriwardana/Book-Management-System-with-Golang-and-Mysql/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRouters = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
