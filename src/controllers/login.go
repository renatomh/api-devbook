package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"api/src/security"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Login is the function which allows user to authenticate and use the API
func Login(w http.ResponseWriter, r *http.Request) {
	// Getting request body
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// If something goes wrong, we call the error response handling function
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	// Initializing the user, reading data from the request body
	var user models.User
	if err = json.Unmarshal(requestBody, &user); err != nil {
		// If something goes wrong, we call the error response handling function
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	// Connecting to the database
	db, err := database.Connect()
	if err != nil {
		// If something goes wrong, we call the error response handling function
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	// Creating the users' repository
	repository := repositories.NewUsersRepository(db)
	// Creating a new user on the repository
	databaseSavedUser, err := repository.SearchByEmail(user.Email)
	if err != nil {
		// If something goes wrong, we call the error response handling function
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	// Checking if password is correct
	if err = security.CheckPassword(user.Pass, databaseSavedUser.Pass); err != nil {
		// If something goes wrong, we call the error response handling function
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	// Generating the user token
	token, _ := authentication.CreateToken(databaseSavedUser.ID)

	w.Write([]byte(token))
}
