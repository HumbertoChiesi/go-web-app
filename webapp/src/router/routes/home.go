package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var homePageRout = Rout{
	URI:                "/home",
	Method:             http.MethodGet,
	Function:           controllers.LoadHomePage,
	NeedAuthentication: true,
}
