package router

import (
	"web-app/src/router/routes"

	"github.com/gorilla/mux"
)

// Gen generater a new router with the configured routes
func Gen() *mux.Router {
	
	return routes.Config(mux.NewRouter())
}
