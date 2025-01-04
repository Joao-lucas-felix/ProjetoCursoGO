package routes

import (
	"net/http"

	"github.com/Joao-lucas-felix/DevBook/API/src/controllers"
)

var PostRoutes = []Route{
	{
		URI: "/post",
		Metodo: http.MethodPost,
		Func: controllers.CreatePost,
		RequireAuth: true,
	},
	{
		URI: "/post",
		Metodo: http.MethodGet,
		Func: controllers.GetAllPosts,
		RequireAuth: true,
	},
	{
		URI: "/post/{postId}",
		Metodo: http.MethodGet,
		Func: controllers.GetPostById,
		RequireAuth: true,
	},
	{
		URI: "/post/{postId}",
		Metodo: http.MethodPut,
		Func: controllers.UpdatePost,
		RequireAuth: true,
	},
	{
		URI: "/post/{postId}",
		Metodo: http.MethodDelete,
		Func: controllers.DeletePost,
		RequireAuth: true,
	},
	{
		URI: "/users/{userId}/posts",
		Metodo: http.MethodGet,
		Func: controllers.GetUserPosts,
		RequireAuth: true,
	},
}	