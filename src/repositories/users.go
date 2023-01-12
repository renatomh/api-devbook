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

// SearchByID a specific user by its ID
func (repository Users) SearchByID(ID uint64) (models.User, error) {
	// Executing the select statement (we won't return the users passwords)
	rows, err := repository.db.Query(
		"select id, name, username, email, createdAt from users where ID = ?",
		ID,
	)
	if err != nil {
		// We return an empty user if an error occurs
		return models.User{}, err
	}
	defer rows.Close()

	// Reading row data
	var user models.User
	if rows.Next() {
		// Getting user
		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Username,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			// We return an empty user if an error occurs
			return models.User{}, err
		}
	}

	// Returning the user data
	return user, nil
}

// Update will edit a specific user data by its ID
func (repository Users) Update(ID uint64, user models.User) error {
	// Preparing the statement to execute the SQL query
	statement, err := repository.db.Prepare(
		"update users set name = ?, username = ?, email = ? where id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	// Executing the update statement
	if _, err = statement.Exec(user.Name, user.Username, user.Email, ID); err != nil {
		return err
	}

	// Returning the function
	return nil
}

// Delete removes a specific user from the database
func (repository Users) Delete(ID uint64) error {
	// Preparing the statement to execute the SQL query
	statement, err := repository.db.Prepare("delete from users where id = ?")
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

// SearchByEmail a specific user by its email, as well as its hashpass (for login purposes)
func (repository Users) SearchByEmail(email string) (models.User, error) {
	// Executing the select statement (we will get only ID, the email and hash password)
	rows, err := repository.db.Query("select id, pass from users where email = ?", email)
	if err != nil {
		// We return an empty user if an error occurs
		return models.User{}, err
	}
	defer rows.Close()

	// Reading row data
	var user models.User
	if rows.Next() {
		// Getting user
		if err = rows.Scan(
			&user.ID,
			&user.Pass,
		); err != nil {
			// We return an empty user if an error occurs
			return models.User{}, err
		}
	}

	// Returning the user data
	return user, nil
}

// Follow allows an user to follow another one
func (repository Users) Follow(userID, followerID uint64) error {
	// Preparing the insert statment
	// We'll ignore the insertion of duplicate entries
	statement, err := repository.db.Prepare(
		"insert ignore into followers (user_id, follower_id) values (?, ?)",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	// Executing the query to follow the user
	if _, err := statement.Exec(userID, followerID); err != nil {
		return err
	}

	// If everything is ok, no error will be returned
	return nil
}

// Unfollow allows an user to stop following another one
func (repository Users) Unfollow(userID, followerID uint64) error {
	// Preparing the delete statment
	statement, err := repository.db.Prepare(
		"delete from followers where user_id = ? and follower_id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	// Executing the query to unfollow the user
	if _, err := statement.Exec(userID, followerID); err != nil {
		return err
	}

	// If everything is ok, no error will be returned
	return nil
}

// SearchFollowers returns an user followers by its ID
func (repository Users) SearchFollowers(userID uint64) ([]models.User, error) {
	// Executing the select statement
	// Here, we're making a join between the users and followers tables
	rows, err := repository.db.Query(`
		select u.id, u.name, u.username, u.email, createdAt
		from users u inner join followers f on u.id = f.follower_id
		where f.user_id = ?`,
		userID)
	if err != nil {
		// We return an empty user if an error occurs
		return nil, err
	}
	defer rows.Close()

	// Reading row data
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

// SearchFollowing returns users followed by another one
func (repository Users) SearchFollowing(userID uint64) ([]models.User, error) {
	// Executing the select statement
	// Here, we're making a join between the users and followers tables
	rows, err := repository.db.Query(`
		select u.id, u.name, u.username, u.email, createdAt
		from users u inner join followers f on u.id = f.user_id
		where f.follower_id = ?`,
		userID)
	if err != nil {
		// We return an empty user if an error occurs
		return nil, err
	}
	defer rows.Close()

	// Reading row data
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
