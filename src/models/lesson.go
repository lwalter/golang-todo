package models

import (
	"time"
)

// Lesson bleh
type Lesson struct {
	ID          int          `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	LessonItems []LessonItem `json:"lessonItems"`
	CreatedAt   time.Time    `json:"createdAt"`
	ModifiedAt  time.Time    `json:"modifiedAt"`
	CreatedBy   string       `json:"createdBy"`
}
