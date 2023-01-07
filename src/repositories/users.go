package repositories

import (
	"api/src/models"
	"database/sql"
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
