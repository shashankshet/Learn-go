package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq" // PostgreSQL driver
)

var DB *sql.DB

// InitDB initializes the database connection
func InitDB() {
	// Fetch credentials from environment variables
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	// Validate environment variables
	if dbUser == "" || dbPassword == "" || dbName == "" || dbHost == "" || dbPort == "" {
		log.Fatal("One or more required environment variables are missing")
	}

	// Connection string
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		dbUser, dbPassword, dbName, dbHost, dbPort)

	var err error
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Failed to connect to Database", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("Failed to ping Database", err)
	}

	log.Println("Database connection established")
}
