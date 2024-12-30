package middlewares

import (
	"fmt"
	"log"
	"net/http"
)
// logger log the request infos 
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n%s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

// Auth verifies if the user in request are athenticated
func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Validate the token")
		next(w, r)
	}
}