package routes

import (
	"api/src/controllers"
	"net/http"
)

var UserRoutes = []Route{
	{
		URI:                   "/users",
		Method:                http.MethodPost,
		Function:              controllers.CreateUser,
		RequireAuthentication: false,
	},
	{
		URI:                   "/users",
		Method:                http.MethodGet,
		Function:              controllers.FindUsers,
		RequireAuthentication: true,
	},
	{
		URI:                   "/users/{userId}",
		Method:                http.MethodGet,
		Function:              controllers.FindUser,
		RequireAuthentication: true,
	},
	{
		URI:                   "/users/{userId}",
		Method:                http.MethodPut,
		Function:              controllers.UpdateUser,
		RequireAuthentication: true,
	},
	{
		URI:                   "/users/{userId}",
		Method:                http.MethodDelete,
		Function:              controllers.DeleteUser,
		RequireAuthentication: true,
	},
	{
		URI:                   "/users/{userId}/follow",
		Method:                http.MethodPost,
		Function:              controllers.FollowUser,
		RequireAuthentication: true,
	},
	{
		URI:                   "/users/{userId}/unfollow",
		Method:                http.MethodPost,
		Function:              controllers.UnFollowUser,
		RequireAuthentication: true,
	},
	{
		URI:                   "/users/{userId}/followers",
		Method:                http.MethodPost,
		Function:              controllers.SearchFollowers,
		RequireAuthentication: true,
	},
	{
		URI:                   "/users/{userId}/following",
		Method:                http.MethodPost,
		Function:              controllers.SearchFollowing,
		RequireAuthentication: true,
	},
	{
		URI:                   "/users/{userId}/update-password",
		Method:                http.MethodPost,
		Function:              controllers.UpdatePassword,
		RequireAuthentication: true,
	},
}
