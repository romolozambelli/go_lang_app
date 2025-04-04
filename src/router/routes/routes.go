package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI          string
	Method       string
	Function     func(http.ResponseWriter, *http.Request)
	Authenticate bool
}

// Setup the routes inside of the router
func Setup(router *mux.Router) *mux.Router {

	routes := routesUser
	routes = append(routes, routesLogin...)

	for _, route := range routes {

		router.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	fileServer := http.FileServer(http.Dir("./assets"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return router
}
