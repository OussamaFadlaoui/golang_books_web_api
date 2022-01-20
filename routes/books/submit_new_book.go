package books

import (
	"fmt"
	"net/http"
)

func SubmitNewBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Body)
}
