package middlewares

import (
	"log"
	"net/http"

	"github.com/Joao-lucas-felix/DevBook/API/src/auth"
	"github.com/Joao-lucas-felix/DevBook/API/src/responses"
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
		if err := auth.ValidateToken(r); err != nil {
			responses.Error(w, http.StatusUnauthorized, err)
			return
		}
		next(w, r)
	}
}
