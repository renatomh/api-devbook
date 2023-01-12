package models

// Password represents an user's password change request format
type Password struct {
	New     string `json:"new"`
	Current string `json:"current"`
}
