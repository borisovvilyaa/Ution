package handlers

import (
	"encoding/json"
	"net/http"
	"ution/internal/models/Errors/database" // Adjust the import path to match your project structure
	"ution/internal/storage"

	"ution/internal/models/User" // Adjust the import path to match your project structure
)

// RegisterUser handles the user registration API request.
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	// Set the content type to application/json
	w.Header().Set("Content-Type", "application/json")

	// Parse the JSON request body
	var user models.User // Assume you have a User struct in your storage package
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	// Call the SignIn function to register the user
	if err := storage.SignIn(user.Username, user.Email, user.PasswordHash, user.FirstName, user.LastName); err != nil {
		if err == database.UserAlreadyExistsError {
			http.Error(w, err.Error(), http.StatusConflict) // 409 Conflict for duplicate users
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError) // 500 Internal Server Error for other errors
		}
		return
	}

	// Respond with a success message
	response := map[string]string{"message": "User registered successfully"}
	w.WriteHeader(http.StatusCreated) // 201 Created
	json.NewEncoder(w).Encode(response)
}
