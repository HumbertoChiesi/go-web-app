package controllers

import (
	dBase "api/src/DB"
	"api/src/authentication"
	"api/src/models"
	repository "api/src/repositories"
	"api/src/responses"
	"api/src/security"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

func Login(w http.ResponseWriter, r *http.Request) {
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

	db, err := dBase.Connect()
	if err != nil {
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.NewUsersRepository(db)
	userSavedInDB, err := repository.SearchByEmail(user.Email)
	if err != nil {
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	}

	if err = security.VerifyPassword(userSavedInDB.Password, user.Password); err != nil {
		responses.ERR(w, http.StatusUnauthorized, err)
		return
	}

	token, err := authentication.CreateToken(userSavedInDB.ID)
	if err != nil {
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	}

	userId := strconv.FormatUint(userSavedInDB.ID, 10)

	responses.JSON(w, http.StatusOK, models.AuthenticationData{ID: userId, Token: token})
}
