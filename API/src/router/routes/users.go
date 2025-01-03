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
	{
		URI:         "/users/{userID}/unfollow",
		Metodo:      http.MethodPost,
		Func:        controllers.UnfollowUser,
		RequireAuth: true,
	},
	{
		URI:         "/users/{userID}/followers",
		Metodo:      http.MethodGet,
		Func:        controllers.GetFollowers,
		RequireAuth: true,
	},
	{
		URI:         "/users/{userID}/following",
		Metodo:      http.MethodGet,
		Func:        controllers.GetFollowings,
		RequireAuth: true,
	},
	{
		URI:         "/users/{userID}/redifine-password",
		Metodo:      http.MethodPost,
		Func:        controllers.RedifinePassword,
		RequireAuth: true,
	},
}
