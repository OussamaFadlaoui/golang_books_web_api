package main

import (
	"github.com/gorilla/mux"
	"golang_books_web_api/middlewares"
	"golang_books_web_api/routes"
	"log"
	"net/http"
)

func handleRequests() {
	router := mux.NewRouter()

	router.NotFoundHandler = http.HandlerFunc(routes.NotFoundPage)

	bookRoutes := router.PathPrefix("/books").Subrouter()
	bookRoutes.HandleFunc("", routes.FetchAllBooks).Methods("GET")
	bookRoutes.HandleFunc("", routes.SubmitNewBook).Methods("POST")

	applyGlobalMiddlewares(router)

	log.Fatal(http.ListenAndServe(":1212", router))
}

func applyGlobalMiddlewares(router *mux.Router) {
	router.Use(
		middlewares.ContentTypeJsonMiddleware,
		middlewares.LoggingMiddleware,
	)
}

func main() {
	handleRequests()
}
