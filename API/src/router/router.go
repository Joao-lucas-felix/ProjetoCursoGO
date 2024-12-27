package router

import (
	"github.com/Joao-lucas-felix/DevBook/API/src/router/routes"
	"github.com/gorilla/mux"
)

// GenRouter return a new Router, with the routes configured
func GenRouter() *mux.Router {
	router := routes.ConfigRoutes(mux.NewRouter())
	return router
}
