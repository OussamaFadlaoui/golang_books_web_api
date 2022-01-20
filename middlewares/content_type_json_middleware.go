package middlewares

import (
	"fmt"
	"net/http"
	"time"
)

func ContentTypeJsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(fmt.Sprintf("[%v] %v", time.Now().Format("2006-01-02 15:04:10"), r.RequestURI))
		next.ServeHTTP(w, r)
	})
}
