package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route represents all the API routes
type Route struct {
	URI                    string
	Method                 string
	Function               func(http.ResponseWriter, *http.Request)
	RequiresAuthentication bool
}

// SetUp adds all routes to the router
func SetUp(r *mux.Router) *mux.Router {
	// Getting the users routes
	routes := usersRoutes

	// For each created route
	for _, route := range routes {
		// Setting the handler function for the route
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	// Returning the new router
	return r
}
