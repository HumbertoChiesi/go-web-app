package routes

import (
	"api/src/controllers"
	"net/http"
)

var postsRout = []Rout{
	{
		URI:                "/posts",
		Method:             http.MethodPost,
		Function:           controllers.CreatePost,
		NeedAuthentication: true,
	},

	{
		URI:                "/posts",
		Method:             http.MethodGet,
		Function:           controllers.SearchPosts,
		NeedAuthentication: true,
	},

	{
		URI:                "/posts/{postId}",
		Method:             http.MethodGet,
		Function:           controllers.SearchPost,
		NeedAuthentication: true,
	},

	{
		URI:                "/posts/{postId}",
		Method:             http.MethodPost,
		Function:           controllers.UpdatePost,
		NeedAuthentication: true,
	},

	{
		URI:                "/posts/{postId}",
		Method:             http.MethodDelete,
		Function:           controllers.DeletePost,
		NeedAuthentication: true,
	},

	{
		URI:                "/users/{userId}/posts",
		Method:             http.MethodGet,
		Function:           controllers.SearchPostByUser,
		NeedAuthentication: true,
	},

	{
		URI:                "/posts/{postId}/like",
		Method:             http.MethodPost,
		Function:           controllers.LikePost,
		NeedAuthentication: true,
	},

	{
		URI:                "/posts/{postId}/unlike",
		Method:             http.MethodPost,
		Function:           controllers.UnlikePost,
		NeedAuthentication: true,
	},
}
