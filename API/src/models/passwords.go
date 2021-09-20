package models

type Password struct {
	Password    string `json:"password"`
	NewPassword string `json:"newPassword"`
}
