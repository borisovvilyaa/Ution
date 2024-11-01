package models

import (
	"time"
)

// User represents a user in the system.
type User struct {
	ID                        int                    `json:"id" db:"id"`
	Username                  string                 `json:"username" db:"username"`
	Email                     string                 `json:"email" db:"email"`
	PasswordHash              string                 `json:"-" db:"password_hash"`
	IsEmailApproved           bool                   `json:"is_email_approved" db:"is_email_approved"`
	CreatedAt                 time.Time              `json:"created_at" db:"created_at"`
	UpdatedAt                 time.Time              `json:"updated_at" db:"updated_at"`
	FirstName                 string                 `json:"first_name" db:"first_name"`
	LastName                  string                 `json:"last_name" db:"last_name"`
	ProfilePicture            string                 `json:"profile_picture" db:"profile_picture"`
	Bio                       string                 `json:"bio" db:"bio"`
	IsActive                  bool                   `json:"is_active" db:"is_active"`
	LastLoginAt               time.Time              `json:"last_login_at" db:"last_login_at"`
	EmailNotificationsEnabled bool                   `json:"email_notifications_enabled" db:"email_notifications_enabled"`
	Settings                  map[string]interface{} `json:"settings" db:"settings"`
}
