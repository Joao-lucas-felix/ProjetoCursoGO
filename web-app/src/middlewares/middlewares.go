package middlewares

import (
	"log"
	"net/http"
	"web-app/src/cookies"
)

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s \n", r.Method, r.URL, r.Host)
		next(w, r)
	}
}

func Auhtenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := cookies.GetToken(r)
		if err != nil {
			http.Redirect(w,r, "/login", http.StatusFound)
		}
		next(w, r)
	}
}
