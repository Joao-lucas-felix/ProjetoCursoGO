package router

import (
	"github.com/Joao-lucas-felix/DevBook/API/src/router/routes"
	"github.com/gorilla/mux"
)

// Gerar return a new Router, with the routes configureds
func GenRouter() *mux.Router {
	router := routes.ConfigRoutes(mux.NewRouter())
	return router
}
