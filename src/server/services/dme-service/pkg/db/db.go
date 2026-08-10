package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var database *sql.DB

// Connect initializes the database connection.
func Connect(databaseURL string) {
	var err error
	database, err = sql.Open("postgres", databaseURL)
	if err != nil {
		log.Fatalf("Failed to open database connection: %v", err)
	}

	// Verify connection
	if err = database.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	log.Println("Database connection successfully established")
}

// GetDB returns the database connection instance.
func GetDB() *sql.DB {
	return database
}
