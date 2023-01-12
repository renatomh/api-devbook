package repositories

import (
	"api/src/models"
	"database/sql"
)

// Posts represents a posts repository
type Posts struct {
	db *sql.DB
}

// NewPostsRepository instantiates/initializes a posts repository
func NewPostsRepository(db *sql.DB) *Posts {
	return &Posts{db}
}

// Create is a Posts' method to create new posts on the repository
func (repository Posts) Create(post models.Post) (uint64, error) {
	// Preparing the insert statment
	statement, err := repository.db.Prepare(
		"insert into posts (title, content, author_id) values(?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	// Executing the query to create new post
	result, err := statement.Exec(post.Title, post.Content, post.AuthorID)
	if err != nil {
		return 0, err
	}

	// Getting the last inserted post ID
	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	// Finally, we return the inserted post ID
	return uint64(lastInsertedId), nil
}
