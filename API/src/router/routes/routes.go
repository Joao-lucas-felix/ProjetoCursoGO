package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route are a struct that represents all routes of API
type Route struct {
	URI         string
	Metodo      string
	Func        func(http.ResponseWriter, *http.Request)
	RequireAuth bool
}

// ConfigRoutes Put all routes in the router
func ConfigRoutes(r *mux.Router) *mux.Router {
	routes := UserRoutes
	routes = append(routes, loginRoute)

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Func).Methods(route.Metodo)
	}

	return r
}
