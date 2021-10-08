package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/requests"
	"webapp/src/responses"
)

//CreatePost calls the API to create a post in the DB
func CreatePost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	post, err := json.Marshal(map[string]string{
		"title":   r.FormValue("title"),
		"content": r.FormValue("content"),
	})

	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrAPI{Err: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/posts", config.APIURL)
	response, err := requests.MakeRequestWithAuthentication(r, http.MethodPost, url, bytes.NewBuffer(post))
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrAPI{Err: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.ProcessStatusCodeErr(w, response)
	}

	responses.JSON(w, response.StatusCode, nil)
}
