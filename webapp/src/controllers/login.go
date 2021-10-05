package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/models"
	"webapp/src/responses"
)

func SignIn(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	user, err := json.Marshal(map[string]string{
		"email":    r.FormValue("email"),
		"password": r.FormValue("password"),
	})

	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrAPI{Err: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/login", config.APIURL)
	response, err := http.Post(url, "application/json", bytes.NewBuffer(user))
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrAPI{Err: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.ProcessStatusCodeErr(w, response)
		return
	}

	var authenticationData models.AuthenticationData
	if err = json.NewDecoder(response.Body).Decode(&authenticationData); err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.ErrAPI{Err: err.Error()})
		return
	}

	if err = cookies.Save(w, authenticationData.ID, authenticationData.Token); err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.ErrAPI{Err: err.Error()})
		return
	}

	responses.JSON(w, http.StatusOK, nil)
}
