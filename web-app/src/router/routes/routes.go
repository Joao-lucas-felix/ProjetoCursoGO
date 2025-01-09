package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Routes are the struct that represents a route in the app
type Routes struct {
	URI         string
	Method      string
	Func         func(http.ResponseWriter, *http.Request)
	RequireAuth bool
}

// Config receives the app router and configs the app routes 
func Config(router *mux.Router) *mux.Router {
	routes := routesLogin

	for _, r := range routes {
		router.HandleFunc(r.URI, r.Func).Methods(r.Method)
	}
	fileServer := http.FileServer(http.Dir("./assets/"))

	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))
	
	return router
 }
