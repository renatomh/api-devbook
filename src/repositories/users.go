package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

// Users represents an users repository
type Users struct {
	db *sql.DB
}

// NewUsersRepository instantiates/initializes a users repository
func NewUsersRepository(db *sql.DB) *Users {
	return &Users{db}
}

// Create is a Users' method to create new users on the repository
func (repository Users) Create(user models.User) (uint64, error) {
	// Preparing the insert statment
	statement, err := repository.db.Prepare(
		"insert into users (name, username, email, pass) values(?, ?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	// Executing the query to create new user
	result, err := statement.Exec(user.Name, user.Username, user.Email, user.Pass)
	if err != nil {
		return 0, err
	}

	// Getting the last inserted user ID
	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	// Finally, we return the inserted user ID
	return uint64(lastInsertedId), nil
}

// Search all users with specified name or username
func (repository Users) Search(nameOrUsername string) ([]models.User, error) {
	// Formatting the query parameter
	nameOrUsername = fmt.Sprintf("%%%s%%", nameOrUsername) // -> %nameOrUsername%

	// Executing the select statement (we won't return the users passwords)
	rows, err := repository.db.Query(
		"select id, name, username, email, createdAt from users where name LIKE ? or username LIKE ?",
		nameOrUsername, nameOrUsername,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Reading rows data
	var users []models.User
	for rows.Next() {
		// Getting user
		var user models.User
		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Username,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}
		// Appending to the users list
		users = append(users, user)
	}

	// Returning the users slice
	return users, nil
}
