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
}
