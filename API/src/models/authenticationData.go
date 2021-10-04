package models

//AuthenticationData contains the authenticated user's token and ID
type AuthenticationData struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}
