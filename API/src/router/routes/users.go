package routes

import (
	"net/http"

	"github.com/Joao-lucas-felix/DevBook/API/src/controllers"
)

// UserRoutes var to define all users routes
var UserRoutes = []Route{
	{
		URI:         "/users",
		Metodo:      http.MethodPost,
		Func:        controllers.CreateUser,
		RequireAuth: false,
	},
	{
		URI:         "/users",
		Metodo:      http.MethodGet,
		Func:        controllers.GetAllUser,
		RequireAuth: true,
	},
	{
		URI:         "/users/{userID}",
		Metodo:      http.MethodGet,
		Func:        controllers.GetUser,
		RequireAuth: true,
	},
	{
		URI:         "/users/{userID}",
		Metodo:      http.MethodPut,
		Func:        controllers.UpdateUser,
		RequireAuth: true,
	},
	{
		URI:         "/users/{userID}",
		Metodo:      http.MethodDelete,
		Func:        controllers.DeleteUser,
		RequireAuth: true,
	},
	{
		URI:         "/users/{userID}/follow",
		Metodo:      http.MethodPost,
		Func:        controllers.FollowUser,
		RequireAuth: true,
	},
}
