package main

import (
	"github.com/gorilla/mux"
	"golang_books_web_api/middlewares"
	"golang_books_web_api/routes"
	"golang_books_web_api/routes/books"
	"log"
	"net/http"
)

func handleRequests() {
	router := mux.NewRouter()

	router.NotFoundHandler = http.HandlerFunc(routes.NotFoundPage)
	router.HandleFunc("/books", books.FetchAllBooks)

	router.Use(middlewares.LoggingMiddleware)

	log.Fatal(http.ListenAndServe(":10000", router))
}

func main() {
	handleRequests()
}
