package data

import (
	"database/sql"
	"log"

	"github.com/lwalter/lessonshare-api/src/models"
)

var db *sql.DB

func init() {
	connStr := "postgres://myapp:dbpass@localhost:5432/myapp?sslmode=disable"
	var err error
	db, err = sql.Open("postgres", connStr)

	if err != nil || db == nil {
		log.Fatal("Unable to connect to db")
		panic("Unable to connect to db")
	}
}

func GetLesson(id int) (*models.Lesson, error) {
	var lesson models.Lesson
	err := db.QueryRow("SELECT * FROM public.lessons WHERE id = $1", id).Scan(&lesson.ID, &lesson.Name)

	if err != nil {
		return nil, err
	}

	return &lesson, nil
}

func GetAllLessons() ([]models.Lesson, error) {
	rows, err := db.Query("SELECT * FROM public.lessons")

	if err != nil {
		// TODO(lnw) how best to handle? log and return empty array?
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
