package routes

import (
	"api/src/controllers"
	"net/http"
)

var loginRout = Rout{
	URI:                "/login",
	Method:             http.MethodPost,
	Function:           controllers.Login,
	NeedAuthentication: false,
}
