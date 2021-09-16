package controllers

import (
	dBase "api/src/DB"
	"api/src/authentication"
	"api/src/models"
	repository "api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

//CreateUser inserts an user in the DB
func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERR(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(requestBody, &user); err != nil {
		responses.ERR(w, http.StatusBadRequest, err)
		return
	}

	if err = user.Prepare("registering"); err != nil {
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	}

	db, err := dBase.Connect()
	if err != nil {
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.NewUsersRepository(db)
	user.ID, err = repository.Create(user)
	if err != nil {
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, user)
}

//SearchUsers searches for all the users in the DB
func SearchUsers(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))

	db, err := dBase.Connect()
	if err != nil {
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.NewUsersRepository(db)
	users, err := repository.Search(nameOrNick)
	if err != nil {
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, users)
}

//SearchUser searches for an specific user in the DB (by the user's ID)
func SearchUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	userID, err := strconv.ParseUint(parameters["userId"], 10, 64)
	if err != nil {
		responses.ERR(w, http.StatusBadRequest, err)
		return
	}

	db, err := dBase.Connect()
	if err != nil {
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.NewUsersRepository(db)
	user, err := repository.SearchByID(userID)
	if err != nil {
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, user)
}

//UpdateUser update an user's content in the DB
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	userID, err := strconv.ParseUint(parameters["userId"], 10, 64)
	if err != nil {
		responses.ERR(w, http.StatusBadRequest, err)
		return
	}

	userTokenID, err := authentication.GetUserID(r)
	if err != nil {
		responses.ERR(w, http.StatusUnauthorized, err)
		return
	}

	if userID != userTokenID {
		responses.ERR(w, http.StatusForbidden, errors.New("not possible to update an user that is not yours"))
		return
	}

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERR(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(requestBody, &user); err != nil {
		responses.ERR(w, http.StatusBadRequest, err)
		return
	}

	if err = user.Prepare("updating"); err != nil {
		responses.ERR(w, http.StatusBadRequest, err)
		return
	}

	db, err := dBase.Connect()
	if err != nil {
		responses.ERR(w, http.StatusBadRequest, err)
		return
	}
	defer db.Close()

	repository := repository.NewUsersRepository(db)
	if err = repository.Update(userID, user); err != nil {
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

//DeleteUser deletes an user from the DB
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	userID, err := strconv.ParseUint(parameters["userId"], 10, 64)
	if err != nil {
		responses.ERR(w, http.StatusBadRequest, err)
		return
	}

	userTokenID, err := authentication.GetUserID(r)
	if err != nil {
		responses.ERR(w, http.StatusUnauthorized, err)
		return
	}

	if userID != userTokenID {
		responses.ERR(w, http.StatusUnauthorized, errors.New("not possible to delete an user that is not yours"))
		return
	}

	db, err := dBase.Connect()
	if err != nil {
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.NewUsersRepository(db)
	if err = repository.Delete(userID); err != nil {
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

//FollowUser allows an user to follow another
func FollowUser(w http.ResponseWriter, r *http.Request) {
	followerId, err := authentication.GetUserID(r)
	if err != nil {
		responses.ERR(w, http.StatusUnauthorized, err)
		return
	}

	parameters := mux.Vars(r)
	userId, err := strconv.ParseUint(parameters["userId"], 10, 64)
	if err != nil {
		responses.ERR(w, http.StatusBadRequest, errors.New("not possible to follow yourself"))
		return
	}

	if followerId == userId {
		responses.ERR(w, http.StatusForbidden, err)
		return
	}

	db, err := dBase.Connect()
	if err != nil {
		responses.ERR(w, http.StatusForbidden, err)
		return
	}
	defer db.Close()

	repository := repository.NewUsersRepository(db)
	if err := repository.Follow(userId, followerId); err != nil {
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}
