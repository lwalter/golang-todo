package models

import (
	"time"
)

// LessonItem
type LessonItem struct {
	ID         int       `json:"id"`
	Content    string    `json:"content"`
	Order      int       `json:"order"`
	CreatedAt  time.Time `json:"createdAt"`
	ModifiedAt time.Time `json:"modifiedAt"`
	CreatedBy  string    `json:"createdBy"`
}
