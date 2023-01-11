package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// CreateUser inserts a new user on the database
func CreateUser(w http.ResponseWriter, r *http.Request) {
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

	// Preparing user for insertion on database
	if err := user.Preare("register"); err != nil {
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
	user.ID, err = repository.Create(user)
	if err != nil {
		// If something goes wrong, we call the error response handling function
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	// If everything is ok
	responses.JSON(w, http.StatusCreated, user)
}

// SearchUsers searchs all users from the database
func SearchUsers(w http.ResponseWriter, r *http.Request) {
	// Getting the name or username to be used while filtering users on database
	nameOrUsername := strings.ToLower(r.URL.Query().Get(("user")))

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
	// Searching users on the repository
	users, err := repository.Search(nameOrUsername)
	if err != nil {
		// If something goes wrong, we call the error response handling function
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	// Returning users response
	responses.JSON(w, http.StatusOK, users)
}

// SearchUser search a specific user from the database
func SearchUser(w http.ResponseWriter, r *http.Request) {
	// Getting the request parameters
	params := mux.Vars(r)

	// Getting the user ID
	userID, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
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
	// Searching user on the repository
	user, err := repository.SearchByID(userID)
	if err != nil {
		// If something goes wrong, we call the error response handling function
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	// Returning user response
	responses.JSON(w, http.StatusOK, user)
}

// UpdateUser updates a specific user on the database
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Getting the request parameters
	params := mux.Vars(r)

	// Getting the user ID
	userID, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

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

	// Preparing user for insertion on database
	if err := user.Preare("edit"); err != nil {
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
	// Updating an existing user on the repository
	if err = repository.Update(userID, user); err != nil {
		// If something goes wrong, we call the error response handling function
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	// If everything is ok
	responses.JSON(w, http.StatusNoContent, nil)
}

// DeletehUser removes a specific user from the database
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Getting the request parameters
	params := mux.Vars(r)

	// Getting the user ID
	userID, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
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
	// Deleting an existing user from the repository
	if err = repository.Delete(userID); err != nil {
		// If something goes wrong, we call the error response handling function
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	// If everything is ok
	responses.JSON(w, http.StatusNoContent, nil)
}
