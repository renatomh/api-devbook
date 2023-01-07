package database

import (
	"api/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // MySQL Driver
)

// Connect establishes a connection to the database
func Connect() (*sql.DB, error) {
	// Connecting to the database
	db, err := sql.Open("mysql", config.DbConnString)

	// If an error occurs during the connection
	if err != nil {
		return nil, err
	}

	// If it wasn't possible to communicate with the database
	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	// Returning the connection
	return db, nil

}
