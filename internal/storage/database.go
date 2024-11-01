package storage

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// DB holds the database connection pool.
var DB *sql.DB

// InitDB initializes the database connection using environment variables.
func InitDB() error {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: could not load .env file, using system environment variables")
	}

	// Retrieve the environment variables
	host := os.Getenv("PG_HOSTNAME")
	port := os.Getenv("PG_PORT")
	user := os.Getenv("PG_USERNAME")
	password := os.Getenv("PG_PASSWORD")
	dbname := os.Getenv("PG_NAME")
	sslmode := os.Getenv("PG_SSLMODE")

	// Define the PostgreSQL connection string
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode)

	// Open a database connection
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		return fmt.Errorf("failed to open database: %v", err)
	}

	// Verify the database connection
	err = DB.Ping()
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	log.Println("Database connected successfully")
	return nil
}
