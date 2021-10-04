package models

//AuthenticationData contains the ID and Token of a authenticated user
type AuthenticationData struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}
