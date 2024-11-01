package database

import "fmt"

// Predefined database error messages
var (
	ErrRequestFailed       = func(err error) error { return fmt.Errorf("database request failed: %w", err) }
	ErrInvalidStatusCode   = func(code int) error { return fmt.Errorf("database invalid status code: %d", code) }
	ErrResponseDecode      = func(err error) error { return fmt.Errorf("database error decoding response: %w", err) }
	UserAlreadyExistsError = fmt.Errorf("user with this username or email already exists")
	ErrApiKeyMissing       = fmt.Errorf("API_TMDB is not set")
)
