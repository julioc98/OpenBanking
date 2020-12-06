package middleware

import (
	"log"
	"net/http"
)

// Logging log all requests URI
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("METHOD:", r.Method, " | ", "PATH: ", r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
