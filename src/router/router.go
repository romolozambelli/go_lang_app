package router

import (
	"webapp/src/router/routes"

	"github.com/gorilla/mux"
)

// Create a mux with all the routes configured
func Generate() *mux.Router {
	r := mux.NewRouter()
	return routes.Setup(r)
}
