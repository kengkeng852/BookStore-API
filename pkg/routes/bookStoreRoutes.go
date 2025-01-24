package routes

import (
	"gihub.com/kengkeng852/bookstore-api/pkg/controllers"
	"github.com/gorilla/mux"
)

var InitBookStoreRouter = func(router *mux.Router) {
	router.HandleFunc("/book", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}",controllers.GetBookByID).Methods("GET")
	router.HandleFunc("/book/{bookId}",controllers.UpdateBookById).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBookById).Methods("DELETE")
}



