package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"api/src/security"
	"encoding/json"
	"errors"
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

	// Getting the user ID provided on the token
	tokenUserID, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	// If user is trying to update another user's data
	if userID != tokenUserID {
		responses.Error(w, http.StatusForbidden, errors.New("You cannot update another user's data"))
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

	// Getting the user ID provided on the token
	tokenUserID, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	// If user is trying to delete another user
	if userID != tokenUserID {
		responses.Error(w, http.StatusForbidden, errors.New("You cannot delete another user's data"))
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

// FollowUser allows an user to follow another one
func FollowUser(w http.ResponseWriter, r *http.Request) {
	// Getting the follower ID provided on the token
	followerID, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	// Getting the request parameters
	params := mux.Vars(r)

	// Getting the user ID
	userID, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	// If user is trying to follow itself
	if userID == followerID {
		responses.Error(w, http.StatusForbidden, errors.New("You cannot follow yourself"))
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
	// Following an existing user on the repository
	if err = repository.Follow(userID, followerID); err != nil {
		// If something goes wrong, we call the error response handling function
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	// If everything is ok
	responses.JSON(w, http.StatusNoContent, nil)
}

// UnfollowUser allows an user to stop following another one
func UnfollowUser(w http.ResponseWriter, r *http.Request) {
	// Getting the follower ID provided on the token
	followerID, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	// Getting the request parameters
	params := mux.Vars(r)

	// Getting the user ID
	userID, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	// If user is trying to unfollow itself
	if userID == followerID {
		responses.Error(w, http.StatusForbidden, errors.New("You cannot unfollow yourself"))
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
	// Unfollowing an existing user on the repository
	if err = repository.Unfollow(userID, followerID); err != nil {
		// If something goes wrong, we call the error response handling function
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	// If everything is ok
	responses.JSON(w, http.StatusNoContent, nil)
}

// SearchFollowers searchs all followers from an user
func SearchFollowers(w http.ResponseWriter, r *http.Request) {
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
	// Searching followers on the repository
	followers, err := repository.SearchFollowers(userID)
	if err != nil {
		// If something goes wrong, we call the error response handling function
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	// Returning followers response
	responses.JSON(w, http.StatusOK, followers)
}

// SearchFollowing searchs all users followed by another one
func SearchFollowing(w http.ResponseWriter, r *http.Request) {
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
	// Searching users on the repository
	users, err := repository.SearchFollowing(userID)
	if err != nil {
		// If something goes wrong, we call the error response handling function
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	// Returning users response
	responses.JSON(w, http.StatusOK, users)
}

// ChangePassword updates a specific user password on the database
func ChangePassword(w http.ResponseWriter, r *http.Request) {
	// Getting the request parameters
	params := mux.Vars(r)

	// Getting the user ID
	userID, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	// Getting the user ID provided on the token
	tokenUserID, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	// If user is trying to change another user's password
	if userID != tokenUserID {
		responses.Error(w, http.StatusForbidden, errors.New("You cannot change another user's password"))
		return
	}

	// Getting request body
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// If something goes wrong, we call the error response handling function
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	// Initializing the password change, reading data from the request body
	var password models.Password
	if err = json.Unmarshal(requestBody, &password); err != nil {
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
	// Checking if current password matches the user password
	databaseHashPassword, err := repository.SearchPassword(userID)
	if err != nil {
		// If something goes wrong, we call the error response handling function
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	if err = security.CheckPassword(password.Current, databaseHashPassword); err != nil {
		// If something goes wrong, we call the error response handling function
		responses.Error(w, http.StatusUnauthorized, errors.New("Incorrect current password"))
		return
	}

	// Creating the new password hash
	hashPassword, err := security.Hash(password.New)
	if err != nil {
		// If something goes wrong, we call the error response handling function
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	// Changing user's password
	if err = repository.ChangePassword(userID, string(hashPassword)); err != nil {
		// If something goes wrong, we call the error response handling function
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	// If everything is ok
	responses.JSON(w, http.StatusNoContent, nil)
}
