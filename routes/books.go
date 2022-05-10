package routes

import (
	"encoding/json"
	"fmt"
	"github.com/bxcodec/faker/v3"
	"golang_books_web_api/models"
	"net/http"
	"os"
)

var books []models.Book

func populateBooks() {
	for i := 0; i < 10; i++ {
		fakerBook := models.Book{}
		err := faker.FakeData(&fakerBook)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		books = append(books, fakerBook)
	}
}

func FetchAllBooks(w http.ResponseWriter, r *http.Request) {
	populateBooks()

	err := json.NewEncoder(w).Encode(books)
	if err != nil {
		_ = json.NewEncoder(w).Encode("Error request")
	}
}

func SubmitNewBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Body)
}
