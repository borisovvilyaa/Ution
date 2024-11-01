package Errors

import "fmt"

// Predefined error messages
var (
	ErrEnvLoad = fmt.Errorf("error loading .env file")
)
