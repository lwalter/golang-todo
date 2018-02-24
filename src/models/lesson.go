package models

import (
	"time"

	"github.com/lib/pq"
)

// Lesson bleh
type Lesson struct {
	ID          int          `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	LessonItems []LessonItem `json:"lessonItems"`
	CreatedAt   time.Time    `json:"createdAt"`
	ModifiedAt  pq.NullTime  `json:"modifiedAt"`
	CreatedBy   string       `json:"createdBy"`
}
