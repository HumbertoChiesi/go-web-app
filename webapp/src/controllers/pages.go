package controllers

import (
	"net/http"
	"webapp/src/utils"
)

//LoadLoginScreen loads login screen
func LoadLoginScreen(w http.ResponseWriter, r *http.Request) {
	utils.ExecTemplate(w, "login.html", nil)
}

//LoadUserRegisterScreen loads user registration screen
func LoadUserRegisterScreen(w http.ResponseWriter, r *http.Request) {
	utils.ExecTemplate(w, "register.html", nil)
}
