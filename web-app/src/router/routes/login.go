package routes

import (
	"net/http"
	"web-app/src/controllers"
)

var routesLogin = []Routes{
	{
		URI:         "/",
		Method:      http.MethodGet,
		Func:        controllers.LoadLoginPage,
		RequireAuth: false,
	},
	{
		URI:         "/login",
		Method:      http.MethodGet,
		Func:        controllers.LoadLoginPage,
		RequireAuth: false,
	},
}
