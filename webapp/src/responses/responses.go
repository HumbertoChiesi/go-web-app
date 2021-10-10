package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

//Err represents the API error response
type ErrAPI struct {
	Err string `json:"err"`
}

//JSON returns a response in JSON format
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if statusCode != http.StatusNoContent {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			log.Fatal(err)
		}
	}
}

//ProcessStatusCodeErr process the request with status code 400 or above
func ProcessStatusCodeErr(w http.ResponseWriter, r *http.Response) {
	var err ErrAPI
	json.NewDecoder(r.Body).Decode(&err)
	JSON(w, r.StatusCode, err)
}
