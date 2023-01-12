package authentication

import (
	// Importing package as an alias
	"api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// CreateToken generates a JSON web token for the user with defined permissions
func CreateToken(userID uint64) (string, error) {
	// Setting user token permissions
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	// Setting the token duration
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["userId"] = userID
	// Creating the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	// Signing the token and returning it
	// The secret must be generated in a safe way and stored on the .env file
	return token.SignedString([]byte(config.SecretKey))
}

// ValidateToken checks if token provided on the request is valid
func ValidateToken(r *http.Request) error {
	// Getting the token string
	tokenString := extractToken(r)
	// Parsing the token string
	token, err := jwt.Parse(tokenString, returnVerificationKey)
	if err != nil {
		return err
	}

	// Checking if required claims are present on token and this is valid
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	// Otherwise, we'll return the error
	return errors.New("Invalid token")
}

// ExtractUserID returns the user ID present on the token
func ExtractUserID(r *http.Request) (uint64, error) {
	// Getting the token string
	tokenString := extractToken(r)
	// Parsing the token string
	token, err := jwt.Parse(tokenString, returnVerificationKey)
	if err != nil {
		return 0, err
	}

	// Checking if required claims, valid token and getting permissions from token
	if permissions, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId, err := strconv.ParseUint(fmt.Sprintf("%.0f", permissions["userId"]), 10, 64)
		if err != nil {
			return 0, err
		}
		return userId, nil
	}

	// Otherwise, we'll return the error
	return 0, errors.New("Invalid token")
}

// extractToken gets the token provided on the request headers
func extractToken(r *http.Request) string {
	// Getting the request token
	token := r.Header.Get("Authorization")

	// Checking if the "Bearer" string was included
	if len(strings.Split(token, " ")) == 2 {
		// Returning the token
		return strings.Split(token, " ")[1]
	}

	// Otherwise, we return an empty string
	return ""
}

// returnVerificationKey will return the token's verification key
func returnVerificationKey(token *jwt.Token) (interface{}, error) {
	// Checking signing method consistency
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signing method! %v", token.Header["alg"])
	}

	// Returning the token's verification key
	return config.SecretKey, nil
}
