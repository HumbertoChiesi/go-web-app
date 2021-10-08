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
}
