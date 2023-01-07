package routes

import (
	"api/src/controllers"
	"net/http"
)

// Defining the users routes
var usersRoutes = []Route{
	{
		URI:                    "/users",
		Method:                 http.MethodPost,
		Function:               controllers.CreateUser,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/users",
		Method:                 http.MethodGet,
		Function:               controllers.SearchUsers,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/users/{userId}",
		Method:                 http.MethodGet,
		Function:               controllers.SearchUser,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/users/{userId}",
		Method:                 http.MethodPut,
		Function:               controllers.UpdateUser,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/users/{userId}",
		Method:                 http.MethodDelete,
		Function:               controllers.DeleteUser,
		RequiresAuthentication: false,
	},
}
