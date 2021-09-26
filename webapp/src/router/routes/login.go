package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var loginRouts = []Rout{
	{
		URI:                "/",
		Method:             http.MethodGet,
		Function:           controllers.LoadLoginScreen,
		NeedAuthentication: false,
	},

	{
		URI:                "/login",
		Method:             http.MethodGet,
		Function:           controllers.LoadLoginScreen,
		NeedAuthentication: false,
	},
}
