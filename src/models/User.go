package models

import (
	"api/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// User represents a social network user
type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Username  string    `json:"username,omitempty"`
	Email     string    `json:"email,omitempty"`
	Pass      string    `json:"pass,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

// Prepare method calls the other methods to adequate user instance for insertion on database
func (user *User) Prepare(action string) error {
	if err := user.validate(action); err != nil {
		return err
	}
	if err := user.format(action); err != nil {
		return err
	}
	return nil
}

// validate checks if user instance is valid
func (user *User) validate(action string) error {
	// If an error is identified
	if user.Name == "" {
		return errors.New("Name is a required field, cannot be left blank")
	}
	if user.Username == "" {
		return errors.New("Username is a required field, cannot be left blank")
	}
	if user.Email == "" {
		return errors.New("Email is a required field, cannot be left blank")
	}
	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("Email format is not valid")
	}
	if action == "register" && user.Pass == "" {
		return errors.New("Pass is a required field, cannot be left blank")
	}

	// If no error is identified
	return nil
}

// format updates user fields, in order to meet the desired format
func (user *User) format(action string) error {
	// Removing trailing/leading spaces
	user.Name = strings.TrimSpace(user.Name)
	user.Username = strings.TrimSpace(user.Username)
	user.Email = strings.TrimSpace(user.Email)

	// When registering a new user
	if action == "register" {
		// Creating the hash for the user password
		hashPass, err := security.Hash(user.Pass)
		if err != nil {
			return err
		}
		// If everything is ok, we'll save the hash pass for the user
		user.Pass = string(hashPass)
	}

	return nil
}
