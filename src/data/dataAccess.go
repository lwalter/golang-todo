package data

import "github.com/lwalter/lessonshare-api/src/models"

func GetLesson(id int) (*models.Lesson, error) {
	var lesson models.Lesson
	err := db.QueryRow("SELECT * FROM public.lessons WHERE id = $1", id).Scan(&lesson.ID, &lesson.Name, &lesson.Description, &lesson.CreatedAt, &lesson.ModifiedAt, &lesson.CreatedBy)

	if err != nil {
		return nil, err
	}

	return &lesson, nil
}

func GetAllLessons() ([]models.Lesson, error) {
	rows, err := db.Query("SELECT * FROM public.lessons")

	if err != nil {
		return nil, err
	}

	lessons := make([]models.Lesson, 0)
	defer rows.Close()
	for rows.Next() {
		var lesson models.Lesson
		if err := rows.Scan(&lesson.ID, &lesson.Name, &lesson.Description, &lesson.CreatedAt, &lesson.ModifiedAt, &lesson.CreatedBy); err != nil {
			return nil, err
		}
		lessons = append(lessons, lesson)
	}
	return lessons, nil
}
