package models

import "time"

// User represents a social network user
type User struct {
	ID        uint64    `json:"id,onitempty"`
	Name      string    `json:"name,onitempty"`
	Username  string    `json:"username,onitempty"`
	Email     string    `json:"email,onitempty"`
	Pass      string    `json:"pass,onitempty"`
	CreatedAt time.Time `json:"createdAt,onitempty"`
}
