package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

// Generate will return a router with set up routes
func Generate() *mux.Router {
	// Initializing the router
	r := mux.NewRouter()

	// Returning the configured router
	return routes.SetUp(r)
}
