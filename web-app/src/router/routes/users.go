package routes

import (
	"net/http"
	"web-app/src/controllers"
)

var UserRoutes = []Routes{
	{
		URI:         "/create-user",
		Method:      http.MethodGet,
		Func:        controllers.LoadCreateUserPage,
		RequireAuth: false,
	},
	{
		URI:         "/users",
		Method:      http.MethodPost,
		Func:        controllers.CreateUser,
		RequireAuth: false,
	},
}
