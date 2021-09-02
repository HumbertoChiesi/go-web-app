package routes

import (
	"api/src/controllers"
	"net/http"
)

var usersRoutes = []Rout{
	{
		URI:                "/users",
		Method:             http.MethodPost,
		Function:           controllers.CreateUser,
		NeedAuthentication: false,
	},
	{
		URI:                "/users",
		Method:             http.MethodGet,
		Function:           controllers.SearchUsers,
		NeedAuthentication: false,
	},

	{
		URI:                "/users/{userId}",
		Method:             http.MethodGet,
		Function:           controllers.SearchUser,
		NeedAuthentication: false,
	},

	{
		URI:                "/users/{userId}",
		Method:             http.MethodPut,
		Function:           controllers.UpdateUser,
		NeedAuthentication: false,
	},

	{
		URI:                "/users/{userId}",
		Method:             http.MethodDelete,
		Function:           controllers.DeleteUser,
		NeedAuthentication: false,
	},
}
