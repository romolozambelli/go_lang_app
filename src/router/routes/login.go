package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var routesLogin = []Route{
	{
		URI:          "/",
		Method:       http.MethodGet,
		Function:     controllers.LoadLoginPage,
		Authenticate: false,
	},
	{
		URI:          "/login",
		Method:       http.MethodGet,
		Function:     controllers.LoadLoginPage,
		Authenticate: false,
	},
}
