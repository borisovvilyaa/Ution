package storage

import (
	"fmt"
	"log"
	"regexp"
	"time"

	"ution/internal/models/Errors/database" // Adjust the import path to match your project structure
)

// SignIn inserts a new user into the database after performing necessary checks.
func SignIn(username, email, passwordHash, firstName, lastName string) error {
	// Validate email format
	if !isValidEmail(email) {
		return fmt.Errorf("invalid email format: %s", email)
	}

	// Check if the user already exists
	if err := checkUserExists(username, email); err != nil {
		return err
	}

	// Define the SQL insert statement
	query := `
		INSERT INTO users (
			username, email, password_hash, created_at, first_name, last_name
		) VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, created_at
	`

	// Get the current time for the created_at field
	createdAt := time.Now()

	// Execute the insert statement
	var id int
	var createdAtDB time.Time
	err := DB.QueryRow(query, username, email, passwordHash, createdAt, firstName, lastName).Scan(&id, &createdAtDB)
	if err != nil {
		// Use predefined error handling
		return database.ErrRequestFailed(err)
	}

	log.Printf("User signed in successfully with ID: %d, Created At: %v", id, createdAtDB)
	return nil
}

// checkUserExists checks if a user with the same username or email already exists.
func checkUserExists(username, email string) error {
	var exists bool

	query := `
		SELECT EXISTS (
			SELECT 1 FROM users WHERE username = $1 OR email = $2
		)
	`
	err := DB.QueryRow(query, username, email).Scan(&exists)
	if err != nil {
		return database.ErrRequestFailed(err)
	}

	if exists {
		return database.UserAlreadyExistsError
	}

	return nil
}

// isValidEmail checks if the provided email is in a valid format.
func isValidEmail(email string) bool {
	// Basic regex for validating email format
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}
