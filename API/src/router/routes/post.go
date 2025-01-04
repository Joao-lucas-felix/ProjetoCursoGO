package routes

import (
	"net/http"

	"github.com/Joao-lucas-felix/DevBook/API/src/controllers"
)

var PostRoutes = []Route{
	{
		URI: "/posts",
		Metodo: http.MethodPost,
		Func: controllers.CreatePost,
		RequireAuth: true,
	},
	{
		URI: "/posts",
		Metodo: http.MethodGet,
		Func: controllers.GetAllPosts,
		RequireAuth: true,
	},
	{
		URI: "/posts/{postId}",
		Metodo: http.MethodGet,
		Func: controllers.GetPostById,
		RequireAuth: true,
	},
	{
		URI: "/posts/{postId}",
		Metodo: http.MethodPut,
		Func: controllers.UpdatePost,
		RequireAuth: true,
	},
	{
		URI: "/posts/{postId}",
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
	{
		URI: "/posts/{postId}/like",
		Metodo: http.MethodPost,
		Func: controllers.LikePost,
		RequireAuth: true,
	},
}	