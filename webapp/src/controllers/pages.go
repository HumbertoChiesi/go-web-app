package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/models"
	"webapp/src/requests"
	"webapp/src/responses"
	"webapp/src/utils"

	"github.com/gorilla/mux"
)

//LoadLoginScreen loads login screen
func LoadLoginScreen(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Read(r)

	if cookie["token"] != "" {
		http.Redirect(w, r, "/home", http.StatusTemporaryRedirect)
	}
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

	cookie, _ := cookies.Read(r)
	userId, _ := strconv.ParseUint(cookie["id"], 10, 64)

	fmt.Println(response.StatusCode, err)
	utils.ExecTemplate(w, "home.html", struct {
		Posts  []models.Post
		UserId uint64
	}{
		Posts:  posts,
		UserId: userId,
	})
}

//LoadPostEditingPage loads the post editing page
func LoadPostEditingPage(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	postId, err := strconv.ParseUint(parameters["postId"], 10, 64)
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrAPI{Err: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/posts/%d", config.APIURL, postId)
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

	var post models.Post
	if err = json.NewDecoder(response.Body).Decode(&post); err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.ErrAPI{Err: err.Error()})
		return
	}

	utils.ExecTemplate(w, "post-update.html", post)
}

//LoadUsersPage loads the pages with the users that are included in the filter
func LoadUsersPage(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))
	url := fmt.Sprintf("%s/users?user=%s", config.APIURL, nameOrNick)

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

	var users []models.User
	if err = json.NewDecoder(response.Body).Decode(&users); err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.ErrAPI{Err: err.Error()})
		return
	}

	utils.ExecTemplate(w, "users.html", users)
}

//LoadUserProfile Load the user's profile page
func LoadUserProfile(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userId, err := strconv.ParseUint(parameters["userId"], 10, 64)
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrAPI{Err: err.Error()})
		return
	}

	cookie, _ := cookies.Read(r)
	userLoggedId, _ := strconv.ParseUint(cookie["id"], 10, 64)

	if userId == userLoggedId {
		http.Redirect(w, r, "/profile", http.StatusTemporaryRedirect)
	}

	user, err := models.SearchCompleteUser(userId, r)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrAPI{Err: err.Error()})
		return
	}

	utils.ExecTemplate(w, "user.html", struct {
		User         models.User
		UserLoggedId uint64
	}{
		User:         user,
		UserLoggedId: userLoggedId,
	})
}

//LoadLoggedUserProfile Load the logged user's profile page
func LoadLoggedUserProfile(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Read(r)
	userId, _ := strconv.ParseUint(cookie["id"], 10, 64)

	user, err := models.SearchCompleteUser(userId, r)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrAPI{Err: err.Error()})
		return
	}

	utils.ExecTemplate(w, "profile.html", user)
}

//LoadUserEditPage loads the edit user's data page
func LoadUserEditPage(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Read(r)
	userId, _ := strconv.ParseUint(cookie["id"], 10, 64)

	chnnl := make(chan models.User)
	go models.SearchUserData(chnnl, userId, r)
	user := <-chnnl

	if user.ID == 0 {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrAPI{Err: "ERROR while searching the user"})
		return
	}

	utils.ExecTemplate(w, "edit-user.html", user)
}

//LoadPasswordEditPage loads the page to change the password
func LoadPasswordEditPage(w http.ResponseWriter, r *http.Request) {
	utils.ExecTemplate(w, "update-password.html", nil)
}
