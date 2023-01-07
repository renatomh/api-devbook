package controllers

import "net/http"

// CreateUser inserts a new user on the database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write(([]byte("Creating user!")))
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
