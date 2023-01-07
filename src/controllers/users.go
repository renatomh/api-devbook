package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// CreateUser inserts a new user on the database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	// Getting request body
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Initializing the user, reading data from the request body
	var user models.User
	if err = json.Unmarshal(requestBody, &user); err != nil {
		log.Fatal(err)
	}

	// Connecting to the database
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	// Creating the users' repository
	repository := repositories.NewUsersRepository(db)
	// Creating a new user on the repository
	userId, err := repository.Create(user)
	if err != nil {
		log.Fatal(err)
	}

	// If everything is ok
	w.Write([]byte(fmt.Sprintf("Inserted ID: %d", userId)))
}

// SearchUsers searchs all users from the database
func SearchUsers(w http.ResponseWriter, r *http.Request) {
	w.Write(([]byte("Searching all users!")))
}

// SearchUser search a specific user from the database
func SearchUser(w http.ResponseWriter, r *http.Request) {
	w.Write(([]byte("Searching an user!")))
}

// UpdateUser updates a specific user on the database
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write(([]byte("Updating user!")))
}

// DeletehUser removes a specific user from the database
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write(([]byte("Deleting user!")))
}
