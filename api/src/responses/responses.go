package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSON returns a JSON response to the request
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatal(err)
	}
}

// ERROR returns an error in JSON format
func Error(w http.ResponseWriter, statusCode int, Error error) {
	JSON(w, statusCode, struct {
		Error string `json:"error"`
	}{
		Error: Error.Error(),
	})
}
