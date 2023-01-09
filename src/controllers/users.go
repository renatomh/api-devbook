package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// CreateUser inserts a new user on the database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	// Getting request body
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// If somethiing goes wrong, we call the error response handling function
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	// Initializing the user, reading data from the request body
	var user models.User
	if err = json.Unmarshal(requestBody, &user); err != nil {
		// If somethiing goes wrong, we call the error response handling function
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	// Preparing user for insertion on database
	if err := user.Preare(); err != nil {
		// If somethiing goes wrong, we call the error response handling function
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	// Connecting to the database
	db, err := database.Connect()
	if err != nil {
		// If somethiing goes wrong, we call the error response handling function
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	// Creating the users' repository
	repository := repositories.NewUsersRepository(db)
	// Creating a new user on the repository
	user.ID, err = repository.Create(user)
	if err != nil {
		// If somethiing goes wrong, we call the error response handling function
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	// If everything is ok
	responses.JSON(w, http.StatusCreated, user)
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
