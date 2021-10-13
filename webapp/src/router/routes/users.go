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

	{
		URI:                "/users",
		Method:             http.MethodPost,
		Function:           controllers.CreateUser,
		NeedAuthentication: false,
	},

	{
		URI:                "/search-users",
		Method:             http.MethodGet,
		Function:           controllers.LoadUsersPage,
		NeedAuthentication: true,
	},

	{
		URI:                "/users/{userId}",
		Method:             http.MethodGet,
		Function:           controllers.LoadUserProfile,
		NeedAuthentication: true,
	},

	{
		URI:                "/users/{userId}/unfollow",
		Method:             http.MethodPost,
		Function:           controllers.UnfollowUser,
		NeedAuthentication: true,
	},

	{
		URI:                "/users/{userId}/follow",
		Method:             http.MethodPost,
		Function:           controllers.FollowUser,
		NeedAuthentication: true,
	},

	{
		URI:                "/profile",
		Method:             http.MethodGet,
		Function:           controllers.LoadLoggedUserProfile,
		NeedAuthentication: true,
	},

	{
		URI:                "/edit-user",
		Method:             http.MethodGet,
		Function:           controllers.LoadUserEditPage,
		NeedAuthentication: true,
	},

	{
		URI:                "/edit-user",
		Method:             http.MethodPut,
		Function:           controllers.EditUser,
		NeedAuthentication: true,
	},

	{
		URI:                "/update-password",
		Method:             http.MethodGet,
		Function:           controllers.LoadPasswordEditPage,
		NeedAuthentication: true,
	},

	{
		URI:                "/update-password",
		Method:             http.MethodPost,
		Function:           controllers.UpdatePassword,
		NeedAuthentication: true,
	},

	{
		URI:                "/delete-user",
		Method:             http.MethodDelete,
		Function:           controllers.DeleteUser,
		NeedAuthentication: true,
	},
}
