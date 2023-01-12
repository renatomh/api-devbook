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
		RequiresAuthentication: true,
	},
	{
		URI:                    "/users/{userId}",
		Method:                 http.MethodGet,
		Function:               controllers.SearchUser,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/users/{userId}",
		Method:                 http.MethodPut,
		Function:               controllers.UpdateUser,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/users/{userId}",
		Method:                 http.MethodDelete,
		Function:               controllers.DeleteUser,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/users/{userId}/follow",
		Method:                 http.MethodPost,
		Function:               controllers.FollowUser,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/users/{userId}/unfollow",
		Method:                 http.MethodDelete,
		Function:               controllers.UnfollowUser,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/users/{userId}/followers",
		Method:                 http.MethodGet,
		Function:               controllers.SearchFollowers,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/users/{userId}/following",
		Method:                 http.MethodGet,
		Function:               controllers.SearchFollowing,
		RequiresAuthentication: true,
	},
}
