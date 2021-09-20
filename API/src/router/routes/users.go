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
		NeedAuthentication: true,
	},

	{
		URI:                "/users/{userId}",
		Method:             http.MethodGet,
		Function:           controllers.SearchUser,
		NeedAuthentication: true,
	},

	{
		URI:                "/users/{userId}",
		Method:             http.MethodPut,
		Function:           controllers.UpdateUser,
		NeedAuthentication: true,
	},

	{
		URI:                "/users/{userId}",
		Method:             http.MethodDelete,
		Function:           controllers.DeleteUser,
		NeedAuthentication: true,
	},

	{
		URI:                "/users/{userId}/follow",
		Method:             http.MethodPost,
		Function:           controllers.FollowUser,
		NeedAuthentication: true,
	},

	{
		URI:                "/users/{userId}/unfollow",
		Method:             http.MethodPost,
		Function:           controllers.UnfollowUser,
		NeedAuthentication: true,
	},

	{
		URI:                "/users/{userId}/followers",
		Method:             http.MethodGet,
		Function:           controllers.SearchFollowers,
		NeedAuthentication: true,
	},

	{
		URI:                "/users/{userId}/following",
		Method:             http.MethodGet,
		Function:           controllers.SearchFollowing,
		NeedAuthentication: true,
	},

	{
		URI:                "/users/{userId}/update-password",
		Method:             http.MethodPost,
		Function:           controllers.UpdatePassword,
		NeedAuthentication: true,
	},
}
