package controllers

import (
	"net/http"
	"webapp/src/cookies"
)

//LogoutUser logout the user from the web app
func LogoutUser(w http.ResponseWriter, r *http.Request) {
	cookies.Delete(w)
	http.Redirect(w, r, "/login", 302)
}
