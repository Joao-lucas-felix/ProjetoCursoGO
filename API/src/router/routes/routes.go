package routes

import (
	"net/http"

	"github.com/Joao-lucas-felix/DevBook/API/src/middlewares"
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
	routes = append(routes, PostRoutes...)

	for _, route := range routes {
		if route.RequireAuth {
			r.HandleFunc(route.URI, 
				middlewares.Logger(middlewares.Auth(route.Func),
				)).Methods(route.Metodo)
		}else{
			r.HandleFunc(route.URI, middlewares.Logger(route.Func)).Methods(route.Metodo)
		}
	}
	return r
}
