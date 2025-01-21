package routes

import (
	"net/http"
	"web-app/src/controllers"
)

var homeRoute = Routes{
	URI: "/home",
	Method: http.MethodGet,
	Func: controllers.LoadHomePage,
	RequireAuth: true,
}
