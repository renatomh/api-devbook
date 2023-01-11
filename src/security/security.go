package security

import "golang.org/x/crypto/bcrypt"

// Hash will create a hash for the provided password
func Hash(pass string) ([]byte, error) {
	// Creating and returning the password hash
	return bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
}

// CheckPassword will verify if a password string matches a hash created for it
func CheckPassword(passString, hashPass string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(passString))
}
