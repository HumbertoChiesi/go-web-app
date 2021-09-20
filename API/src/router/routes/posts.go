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
}
