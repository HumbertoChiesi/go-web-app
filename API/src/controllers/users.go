package controllers

import (
	db "api/src/DB"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
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

	if err = user.Prepare(); err != nil {
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	}

	db, err := db.Connect()
	if err != nil {
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	user.ID, err = repository.Create(user)
	if err != nil {
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, user)
}

//SearchUsers searches for all the users in the DB
func SearchUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Searching all the users!!"))
}

//SearchUser searches for an specific user in the DB (by the user's ID)
func SearchUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Searching an user!!"))
}

//UpdateUser update an user's content in the DB
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Updating an user!!"))
}

//DeleteUser deletes an user from the DB
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleting an user!!"))
}
