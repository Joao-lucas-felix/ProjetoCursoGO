package routes

import (
	"net/http"

	"github.com/Joao-lucas-felix/DevBook/API/src/controllers"
)

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
		RequireAuth: false,
	},
	{
		URI:         "/users/{userID}",
		Metodo:      http.MethodGet,
		Func:        controllers.GetUser,
		RequireAuth: false,
	},
	{
		URI:         "/users/{userID}",
		Metodo:      http.MethodPut,
		Func:        controllers.UpdateUser,
		RequireAuth: false,
	},
	{
		URI:         "/users/{userID}",
		Metodo:      http.MethodDelete,
		Func:        controllers.DeleteUser,
		RequireAuth: false,
	},
}
