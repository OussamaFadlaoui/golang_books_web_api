package routes

import (
	"encoding/json"
	"golang_books_web_api/structs"
	"net/http"
)

func NotFoundPage(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(structs.JsonResponse{
		Message: "Page not found",
	})
}
