package authentication

import (
	// Importing package as an alias
	"api/src/config"
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
