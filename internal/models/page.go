package models

import (
	"time"
)

type Page struct {
	ID          int       `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	Content     string    `json:"content" db:"content"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	OwnerID     int       `json:"owner_id" db:"owner_id"`
	IsPublic    bool      `json:"is_public" db:"is_public"`
	Tags        []string  `json:"tags" db:"tags"`
	ParentID    *int      `json:"parent_id,omitempty" db:"parent_id"`
	RevisionID  int       `json:"revision_id" db:"revision_id"`
	SharingLink string    `json:"sharing_link,omitempty" db:"sharing_link"`
	Views       int       `json:"views" db:"views"`
	Starred     bool      `json:"starred" db:"starred"`
}
