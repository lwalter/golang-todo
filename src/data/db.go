package data

import "github.com/lwalter/lessonshare-api/src/models"

var lessons []models.Lesson

func init() {
	lessons = append(lessons, models.Lesson{ID: 1, Name: "World War II"})
}

func GetLesson(id int) (*models.Lesson, error) {
	var lesson *models.Lesson
	for _, item := range lessons {
		if item.ID == id {
			lesson = &item
			break
		}
	}

	return lesson, nil
}

func GetAllLessons() ([]models.Lesson, error) {
	return lessons, nil
}
