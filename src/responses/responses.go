package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSON returns a JSON response to the request
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	// Setting the content type
	w.Header().Add("Content-Type", "application/json")
	// Setting the response status code
	w.WriteHeader(statusCode)

	// Returning data as JSON
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatal(err)
	}
}

// Error returns a JSON formatted error response message to the request
func Error(w http.ResponseWriter, statusCode int, err error) {
	// Formats data and call the JSOn function
	JSON(w, statusCode, struct {
		Error string `json:"error"`
	}{
		Error: err.Error(),
	})
}
