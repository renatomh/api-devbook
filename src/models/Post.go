package models

import (
	"errors"
	"strings"
	"time"
)

// Post represents a social network user post
type Post struct {
	ID             uint64    `json:"id,omitempty"`
	Title          string    `json:"title,omitempty"`
	Content        string    `json:"content,omitempty"`
	AuthorID       uint64    `json:"authorId,omitempty"`
	AuthorUsername string    `json:"authorUsername,omitempty"`
	Likes          uint64    `json:"likes"`
	CreatedAt      time.Time `json:"createdAt,omitempty"`
}

// Prepare method calls the other methods to adequate post instance for insertion on database
func (post *Post) Prepare() error {
	if err := post.validate(); err != nil {
		return err
	}
	post.format()
	return nil
}

// validate checks if post instance is valid
func (post *Post) validate() error {
	// If an error is identified
	if post.Title == "" {
		return errors.New("Title is a required field, cannot be left blank")
	}
	if post.Content == "" {
		return errors.New("Content is a required field, cannot be left blank")
	}

	// If no error is identified
	return nil
}

// format updates post fields, in order to meet the desired format
func (post *Post) format() {
	// Removing trailing/leading spaces
	post.Title = strings.TrimSpace(post.Title)
	post.Content = strings.TrimSpace(post.Content)
}
