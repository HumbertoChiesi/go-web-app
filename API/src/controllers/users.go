package controllers

import (
	dBase "api/src/DB"
	"api/src/authentication"
	"api/src/models"
	repository "api/src/repositories"
	"api/src/responses"
	"api/src/security"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

//CreateUser inserts a user in the DB
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

//SearchUser searches for a specific user in the DB (by the user's ID)
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

//UpdateUser update a user's content in the DB
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
		responses.ERR(w, http.StatusForbidden, errors.New("not possible to update a user that is not yours"))
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

//DeleteUser deletes a user from the DB
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
		responses.ERR(w, http.StatusUnauthorized, errors.New("not possible to delete a user that is not yours"))
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

//FollowUser allows a user to follow aother
func FollowUser(w http.ResponseWriter, r *http.Request) {
	followerId, err := authentication.GetUserID(r)
	if err != nil {
		responses.ERR(w, http.StatusUnauthorized, err)
		return
	}

	parameters := mux.Vars(r)
	userId, err := strconv.ParseUint(parameters["userId"], 10, 64)
	if err != nil {
		responses.ERR(w, http.StatusBadRequest, err)
		return
	}

	if followerId == userId {
		responses.ERR(w, http.StatusForbidden, errors.New("not possible to follow yourself"))
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

//UnfollowUser allows a user to unfollow aother
func UnfollowUser(w http.ResponseWriter, r *http.Request) {
	followerId, err := authentication.GetUserID(r)
	if err != nil {
		responses.ERR(w, http.StatusUnauthorized, err)
		return
	}

	parameters := mux.Vars(r)
	userId, err := strconv.ParseUint(parameters["userId"], 10, 64)
	if err != nil {
		responses.ERR(w, http.StatusBadRequest, err)
		return
	}

	if followerId == userId {
		responses.ERR(w, http.StatusForbidden, errors.New("not possible to unfollow yourself"))
		return
	}

	db, err := dBase.Connect()
	if err != nil {
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.NewUsersRepository(db)
	if err = repository.Unfollow(userId, followerId); err != nil {
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

//SearchFollowers searchs a user's followers
func SearchFollowers(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userId, err := strconv.ParseUint(parameters["userId"], 10, 64)
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
	followers, err := repository.SearchFollowers(userId)
	if err != nil {
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, followers)
}

//SearchFollowing searchs the users tha a user is following
func SearchFollowing(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userId, err := strconv.ParseUint(parameters["userId"], 10, 64)
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
	users, err := repository.SearchFollowing(userId)
	if err != nil {
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, users)
}

//UpdatePassword updates a user's password
func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	userTokenId, err := authentication.GetUserID(r)
	if err != nil {
		responses.ERR(w, http.StatusUnauthorized, err)
		return
	}

	parameters := mux.Vars(r)
	userId, err := strconv.ParseUint(parameters["userId"], 10, 64)
	if err != nil {
		responses.ERR(w, http.StatusBadRequest, err)
		return
	}

	if userId != userTokenId {
		responses.ERR(w, http.StatusForbidden, errors.New("not possible to update a user that is not yours"))
		return
	}

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERR(w, http.StatusUnauthorized, err)
		return
	}

	var password models.Password
	if err = json.Unmarshal(requestBody, &password); err != nil {
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
	passwordDB, err := repository.SearchPassword(userId)
	if err != nil {
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	}

	if err = security.VerifyPassword(passwordDB, password.Password); err != nil {
		responses.ERR(w, http.StatusUnauthorized, errors.New("wrong password"))
		return
	}

	hashPassword, err := security.Hash(password.NewPassword)
	if err != nil {
		responses.ERR(w, http.StatusBadRequest, err)
		return
	}

	if err = repository.UpdatePassword(userId, string(hashPassword)); err != nil {
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, nil)
}
