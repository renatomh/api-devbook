package routes

import (
	"api/src/middlewares"
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
	// Getting login route
	routes = append(routes, loginRoute)
	// Getting posts routes
	routes = append(routes, postsRoutes...)

	// For each created route
	for _, route := range routes {
		// If route requires authentication
		if route.RequiresAuthentication {
			// Setting the handler function for the route, using the logger and authentication middlewares
			r.HandleFunc(route.URI,
				middlewares.Logger(middlewares.Authenticate(route.Function)),
			).Methods(route.Method)
		} else {
			// Setting the handler function for the route, using the logger middleware
			r.HandleFunc(route.URI,
				middlewares.Logger(route.Function),
			).Methods(route.Method)
		}
	}

	// Returning the new router
	return r
}
