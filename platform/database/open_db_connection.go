package database

import (
	"os"
)

func OpenDBConnection() {

	// Get DB_TYPE value from .env file.
	dbType := os.Getenv("DB_TYPE")

	// Define a new Database connection with right DB type.
	switch dbType {
	case "pgx":
		PostgreSQLConnection()
	case "mysql":
		return
	default:
		return
	}
}
