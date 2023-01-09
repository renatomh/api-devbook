package models

import (
	"errors"
	"strings"
	"time"
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

// Prepare method calls the other methods to adeuqate user instance for insertion on database
func (user *User) Preare() error {
	if err := user.validate(); err != nil {
		return err
	}
	user.format()
	return nil
}

// validate checks if user instance is valid
func (user *User) validate() error {
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
	if user.Pass == "" {
		return errors.New("Pass is a required field, cannot be left blank")
	}

	// If no error is identified
	return nil
}

// format updates user fields, in order to meet the desired format
func (user *User) format() {
	// Removing trailing/leading spaces
	user.Name = strings.TrimSpace(user.Name)
	user.Username = strings.TrimSpace(user.Username)
	user.Email = strings.TrimSpace(user.Email)
}
