package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var usersRouts = []Rout{
	{
		URI:                "/create-user",
		Method:             http.MethodGet,
		Function:           controllers.LoadUserRegisterScreen,
		NeedAuthentication: false,
	},
}
