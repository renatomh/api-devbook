package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Defining vars to be used by the application
var (
	// Connection string with MySQL
	DbConnString = ""
	// Port number where API will be running
	Port = 0
	// Secret key for JWT signing
	SecretKey []byte
)

// Load initializes environment variables
func Load() {
	var err error
	// Reading .env data
	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	// Setting vars with received data

	// Port number must be converted to int
	Port, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		// Default port number
		Port = 9000
	}

	// Creating database connection string
	DbConnString = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
	)

	// Setting the secret key
	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}
