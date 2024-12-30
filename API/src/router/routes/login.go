package routes

import (
	"net/http"

	"github.com/Joao-lucas-felix/DevBook/API/src/controllers"
)

var loginRoute = Route{
	
	URI: "/login",
	Metodo: http.MethodPost,
	Func: controllers.Login,
	RequireAuth: false,

}