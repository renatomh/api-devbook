package middlewares

import (
	"api/src/authentication"
	"api/src/responses"
	"fmt"
	"net/http"
)

// Logger writes request data on terminal
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		// Goes to the next middleware/request handler function
		next(w, r)
	}
}

// Authenticate checks if user making the request is authenticated
func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Checking if token is valid
		if err := authentication.ValidateToken(r); err != nil {
			responses.Error(w, http.StatusUnauthorized, err)
			return
		}
		// Goes to the next middleware/request handler function
		next(w, r)
	}
}
