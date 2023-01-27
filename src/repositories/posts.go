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
		"insert into posts (title, content, author_id) values (?, ?, ?)",
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

// SearchByID a specific post by its ID
func (repository Posts) SearchByID(postID uint64) (models.Post, error) {
	// Executing the select statement
	rows, err := repository.db.Query(
		`select p.*, u.username from
		posts p inner join users u
		on u.id = p.author_id
		where p.ID = ?`,
		postID,
	)
	if err != nil {
		// We return an empty post if an error occurs
		return models.Post{}, err
	}
	defer rows.Close()

	// Reading row data
	var post models.Post
	if rows.Next() {
		// Getting post
		if err = rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorUsername,
		); err != nil {
			// We return an empty post if an error occurs
			return models.Post{}, err
		}
	}

	// Returning the post data
	return post, nil
}
