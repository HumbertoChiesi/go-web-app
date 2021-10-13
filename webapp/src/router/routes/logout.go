package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var logoutRout = Rout{
	URI:                "/logout",
	Method:             http.MethodGet,
	Function:           controllers.LogoutUser,
	NeedAuthentication: true,
}
