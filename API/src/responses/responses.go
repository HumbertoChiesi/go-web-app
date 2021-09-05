package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

//JSON returns a response in JSON to the requisition
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			log.Fatal(err)
		}
	}
}

//ERR returns an error in JSON format
func ERR(w http.ResponseWriter, statusCode int, err error) {
	JSON(w, statusCode, struct {
		Erro string `json:"err"`
	}{
		Erro: err.Error(),
	})
}
