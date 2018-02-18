package models

import (
	"time"
)

// Lesson bleh
type Lesson struct {
	ID          int       `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
	CreatedAt   time.Time `json:"createdAt,omitempty"`
	ModifiedAt  time.Time `json:"modifiedAt,omitempty"`
	CreatedBy   string    `json:"createdBy,omitempty"`
}
