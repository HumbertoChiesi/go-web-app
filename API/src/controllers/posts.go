package controllers

import (
	dBase "api/src/DB"
	"api/src/authentication"
	"api/src/models"
	repository "api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//CreatePost adds a new post in the DB
func CreatePost(w http.ResponseWriter, r *http.Request) {
	userId, err := authentication.GetUserID(r)
	if err != nil {
		responses.ERR(w, http.StatusUnauthorized, err)
		return
	}

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERR(w, http.StatusUnprocessableEntity, err)
		return
	}

	var post models.Post
	if err = json.Unmarshal(requestBody, &post); err != nil {
		responses.ERR(w, http.StatusBadRequest, err)
		return
	}
	post.PosterID = userId

	db, err := dBase.Connect()
	if err != nil {
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.NewPostRepository(db)
	post.ID, err = repository.Create(post)
	if err != nil {
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, post)
}

//SearchPosts gets the posts that would appear a user's feed
func SearchPosts(w http.ResponseWriter, r *http.Request) {
	userId, err := authentication.GetUserID(r)
	if err != nil {
		responses.JSON(w, http.StatusUnauthorized, err)
		return
	}

	db, err := dBase.Connect()
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.NewPostRepository(db)
	posts, err := repository.Search(userId)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, posts)
}

//SearchPost search a specific post
func SearchPost(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	postId, err := strconv.ParseUint(parameters["postId"], 10, 64)
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, err)
		return
	}

	db, err := dBase.Connect()
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.NewPostRepository(db)
	post, err := repository.SearchById(postId)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, post)
}

//UpdatePost update the data of one post
func UpdatePost(w http.ResponseWriter, r *http.Request) {

}

//DeletePost deletes the data from a specific post
func DeletePost(w http.ResponseWriter, r *http.Request) {

}
