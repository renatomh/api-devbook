package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// CreatePost inserts a new post on the database
func CreatePost(w http.ResponseWriter, r *http.Request) {
	// Getting the user ID provided on the token
	tokenUserID, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	// Getting request body
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// If something goes wrong, we call the error response handling function
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	// Initializing the post, reading data from the request body
	var post models.Post
	if err = json.Unmarshal(requestBody, &post); err != nil {
		// If something goes wrong, we call the error response handling function
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	// Setting the user ID as the post author ID
	post.AuthorID = tokenUserID

	// Preparing post for insertion on database
	if err := post.Prepare(); err != nil {
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

	// Creating the posts' repository
	repository := repositories.NewPostsRepository(db)
	// Creating a new post on the repository
	post.ID, err = repository.Create(post)
	if err != nil {
		// If something goes wrong, we call the error response handling function
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	// If everything is ok
	responses.JSON(w, http.StatusCreated, post)
}

// SearchPosts searchs users and following users posts (user's feed)
func SearchPosts(w http.ResponseWriter, r *http.Request) {
	// Getting the user ID provided on the token
	tokenUserID, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
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

	// Creating the posts' repository
	repository := repositories.NewPostsRepository(db)
	// Searching posts on the repository
	posts, err := repository.Search(tokenUserID)
	if err != nil {
		// If something goes wrong, we call the error response handling function
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	// Returning posts response
	responses.JSON(w, http.StatusOK, posts)
}

// SearchPost search a specific post from the database
func SearchPost(w http.ResponseWriter, r *http.Request) {
	// Getting the request parameters
	params := mux.Vars(r)

	// Getting the post ID
	postID, err := strconv.ParseUint(params["postId"], 10, 64)
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

	// Creating the posts' repository
	repository := repositories.NewPostsRepository(db)
	// Searching post on the repository
	post, err := repository.SearchByID(postID)
	if err != nil {
		// If something goes wrong, we call the error response handling function
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	// Returning post response
	responses.JSON(w, http.StatusOK, post)
}

// UpdatePost updates a specific post on the database
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	// Getting the request parameters
	params := mux.Vars(r)

	// Getting the post ID
	postID, err := strconv.ParseUint(params["postId"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	// Getting the user ID provided on the token
	userID, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
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

	// Creating the posts' repository
	repository := repositories.NewPostsRepository(db)

	// Getting the post saved on the databse by the ID provided
	savedPost, err := repository.SearchByID(postID)
	if err != nil {
		// If something goes wrong, we call the error response handling function
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	// If user is trying to update another user's post
	if savedPost.AuthorID != userID {
		responses.Error(w, http.StatusForbidden, errors.New("You cannot update another user's post"))
		return
	}

	// Getting request body
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// If something goes wrong, we call the error response handling function
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	// Initializing the post, reading data from the request body
	var post models.Post
	if err = json.Unmarshal(requestBody, &post); err != nil {
		// If something goes wrong, we call the error response handling function
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	// Preparing post for update on database
	if err := post.Prepare(); err != nil {
		// If something goes wrong, we call the error response handling function
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	// Updating the existing post on the repository
	if err = repository.Update(postID, post); err != nil {
		// If something goes wrong, we call the error response handling function
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	// If everything is ok
	responses.JSON(w, http.StatusNoContent, nil)
}

// DeletePost removes a specific post from the database
func DeletePost(w http.ResponseWriter, r *http.Request) {
	// Getting the request parameters
	params := mux.Vars(r)

	// Getting the post ID
	postID, err := strconv.ParseUint(params["postId"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	// Getting the user ID provided on the token
	userID, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
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

	// Creating the posts' repository
	repository := repositories.NewPostsRepository(db)

	// Getting the post saved on the databse by the ID provided
	savedPost, err := repository.SearchByID(postID)
	if err != nil {
		// If something goes wrong, we call the error response handling function
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	// If user is trying to delete another user's post
	if savedPost.AuthorID != userID {
		responses.Error(w, http.StatusForbidden, errors.New("You cannot delete another user's post"))
		return
	}

	// Deleting the existing post on the repository
	if err = repository.Delete(postID); err != nil {
		// If something goes wrong, we call the error response handling function
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	// If everything is ok
	responses.JSON(w, http.StatusNoContent, nil)
}

// SearchPostsByUser searchs a specific user posts
func SearchPostsByUser(w http.ResponseWriter, r *http.Request) {
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

	// Creating the posts' repository
	repository := repositories.NewPostsRepository(db)
	// Searching posts on the repository
	posts, err := repository.SearchByUser(userID)
	if err != nil {
		// If something goes wrong, we call the error response handling function
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	// Returning posts response
	responses.JSON(w, http.StatusOK, posts)
}
