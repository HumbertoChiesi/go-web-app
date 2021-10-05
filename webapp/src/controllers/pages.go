package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/models"
	"webapp/src/requests"
	"webapp/src/responses"
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

//LoadHomePage loads the home page with the posts
func LoadHomePage(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/posts", config.APIURL)
	response, err := requests.MakeRequestWithAuthentication(r, http.MethodGet, url, nil)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrAPI{Err: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.ProcessStatusCodeErr(w, response)
		return
	}

	var posts []models.Post
	if err = json.NewDecoder(response.Body).Decode(&posts); err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.ErrAPI{Err: err.Error()})
		return
	}

	fmt.Println(response.StatusCode, err)
	utils.ExecTemplate(w, "home.html", posts)
}
