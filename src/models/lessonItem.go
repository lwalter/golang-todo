package models

import (
	"time"

	"github.com/lib/pq"
)

// LessonItem data entity
type LessonItem struct {
	ID         int         `json:"id"`
	Content    string      `json:"content"`
	Ordinal    int         `json:"order"`
	CreatedAt  time.Time   `json:"createdAt"`
	ModifiedAt pq.NullTime `json:"modifiedAt"`
	CreatedBy  string      `json:"createdBy"`
}
