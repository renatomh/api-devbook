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

// Search posts from user and users followed by the user
func (repository Posts) Search(userID uint64) ([]models.Post, error) {
	// Executing the select statement
	rows, err := repository.db.Query(
		`select distinct p.*, u.username from posts p
		inner join users u on u.id = p.author_id
		inner join followers f on p.author_id = f.user_id
		where p.author_id = ? or f.follower_id = ?
		order by 1 desc;`,
		userID, userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Reading rows data
	var posts []models.Post
	for rows.Next() {
		// Getting post
		var post models.Post
		if err = rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorUsername,
		); err != nil {
			return nil, err
		}
		// Appending to the posts list
		posts = append(posts, post)
	}

	// Returning the posts slice
	return posts, nil
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

// Update will edit a specific post data by its ID
func (repository Posts) Update(ID uint64, post models.Post) error {
	// Preparing the statement to execute the SQL query
	statement, err := repository.db.Prepare(
		"update posts set title = ?, content = ? where id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	// Executing the update statement
	if _, err = statement.Exec(post.Title, post.Content, ID); err != nil {
		return err
	}

	// Returning the function
	return nil
}

// Delete removes a specific post from the database
func (repository Posts) Delete(ID uint64) error {
	// Preparing the statement to execute the SQL query
	statement, err := repository.db.Prepare("delete from posts where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	// Executing the delete statement
	if _, err = statement.Exec(ID); err != nil {
		return err
	}

	// Returning the function
	return nil
}

// SearchByUser returns a specific user posts
func (repository Posts) SearchByUser(userID uint64) ([]models.Post, error) {
	// Executing the select statement
	rows, err := repository.db.Query(
		`select p.*, u.username from posts p
		inner join users u on u.id = p.author_id
		where p.author_id = ?`,
		userID,
	)
	if err != nil {
		// We return an empty list if an error occurs
		return nil, err
	}
	defer rows.Close()

	// Reading rows data
	var posts []models.Post
	for rows.Next() {
		// Getting post
		var post models.Post
		if err = rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorUsername,
		); err != nil {
			return nil, err
		}
		// Appending to the posts list
		posts = append(posts, post)
	}

	// Returning the posts slice
	return posts, nil
}

// Like will add 1 to the number of likes in a post
func (repository Posts) Like(postID uint64) error {
	// Preparing the statement to execute the SQL query
	statement, err := repository.db.Prepare(
		"update posts set likes = likes + 1 where id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	// Executing the update statement
	if _, err = statement.Exec(postID); err != nil {
		return err
	}

	// Returning the function
	return nil
}
