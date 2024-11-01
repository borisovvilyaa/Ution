package models

import (
	"time"
)

type Workspace struct {
	ID          int                    `json:"id" db:"id"`
	Name        string                 `json:"name" db:"name"`
	Description string                 `json:"description" db:"description"`
	OwnerID     int                    `json:"owner_id" db:"owner_id"`
	CreatedAt   time.Time              `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time              `json:"updated_at" db:"updated_at"`
	IsPublic    bool                   `json:"is_public" db:"is_public"`
	Members     []int                  `json:"members" db:"members"`
	Settings    map[string]interface{} `json:"settings" db:"settings"`
}
