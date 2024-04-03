package routes

import (
	"api/src/controllers"
	"net/http"
)

var PostsRoutes = []Route{
	{
		URI:                   "/posts",
		Method:                http.MethodPost,
		Function:              controllers.CreatePost,
		RequireAuthentication: false,
	},
	{
		URI:                   "/posts",
		Method:                http.MethodGet,
		Function:              controllers.FindPosts,
		RequireAuthentication: true,
	},
	{
		URI:                   "/posts/{postId}",
		Method:                http.MethodGet,
		Function:              controllers.FindPost,
		RequireAuthentication: true,
	},
	{
		URI:                   "/posts/{postId}",
		Method:                http.MethodPut,
		Function:              controllers.UpdatePost,
		RequireAuthentication: true,
	},
	{
		URI:                   "/posts/{postId}",
		Method:                http.MethodDelete,
		Function:              controllers.DeletePost,
		RequireAuthentication: true,
	},
	{
		URI:                   "/users/{userId}/posts",
		Method:                http.MethodGet,
		Function:              controllers.SeachPostsByUser,
		RequireAuthentication: true,
	},
	{
		URI:                   "/posts/{postId}/like",
		Method:                http.MethodGet,
		Function:              controllers.LikePost,
		RequireAuthentication: true,
	},
	{
		URI:                   "/posts/{postId}/unlike",
		Method:                http.MethodGet,
		Function:              controllers.UnLikePost,
		RequireAuthentication: true,
	},
}
