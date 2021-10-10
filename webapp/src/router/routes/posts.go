package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var postsRout = []Rout{
	{
		URI:                "/posts",
		Method:             http.MethodPost,
		Function:           controllers.CreatePost,
		NeedAuthentication: true,
	},
	{
		URI:                "/posts/{postId}/like",
		Method:             http.MethodPost,
		Function:           controllers.LikePost,
		NeedAuthentication: true,
	},
	{
		URI:                "/posts/{postId}/dislike",
		Method:             http.MethodPost,
		Function:           controllers.DislikePost,
		NeedAuthentication: true,
	},
	{
		URI:                "/posts/{postId}/edit",
		Method:             http.MethodGet,
		Function:           controllers.LoadPostEditingPage,
		NeedAuthentication: true,
	},
	{
		URI:                "/posts/{postId}",
		Method:             http.MethodPut,
		Function:           controllers.UpdatePost,
		NeedAuthentication: true,
	},
}
