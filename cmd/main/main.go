package main

import (
	"log"
	"net/http"

	"gihub.com/kengkeng852/bookstore-api/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	routes.InitBookStoreRouter(router)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}