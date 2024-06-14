package storage

import (
	"database/sql"
	"mymode/model"
	"time"
)

type NewLessonRepo struct {
	Db *sql.DB
}

func NewLesson(db *sql.DB) *NewLessonRepo {
	return &NewLessonRepo{Db: db}
}

func (L *NewLessonRepo) CreateLesson(lesson model.Lesson) error {
	_, err := L.Db.Exec(`INSERT INTO 
						Lessons(course_id, title, content) 
						VALUES($1, $2, $3)`, lesson.CourseId, lesson.Title, lesson.Content)
	return err
}

func (L *NewLessonRepo) ReadLesson(id string) (model.Lesson, error) {
	lesson := model.Lesson{}
	err := L.Db.QueryRow(`SELECT 
							  title, content FROM Lessons
						  WHERE 
							  lesson_id = $1`, id).Scan(&lesson.Title, &lesson.Content)
	return lesson, err
}

func (L *NewLessonRepo) UpdateLesson(lesson model.Lesson, id string) error {
	_, err := L.Db.Exec(`UPDATE Lessons SET 
							title = $1, content = $2
						WHERE 
							deleted_at is null and lesson_id = $3`,
		lesson.Title, lesson.Content, id)
	return err
}

func (L *NewLessonRepo) DeleteLesson(id string) error {
	_, err := L.Db.Exec(`UPDATE Lessons SET 
							deleted_at = $1
						WHERE 
							deleted_at is null and lesson_id = $2`, time.Now(), id)
	return err
}

func (L *NewLessonRepo) GetAllLessons(filter string, arr []interface{}) ([]model.Lesson, error) {
	lessons := []model.Lesson{}
	rows, err := L.Db.Query(`SELECT 
								lesson_id, course_id, title, content, created_at
							FROM 
								Lessons
							WHERE
								deleted_at is null` + filter, arr...)
	if err != nil {
		return lessons, err
	}

	for rows.Next() {
		lesson := model.Lesson{}
		err = rows.Scan(&lesson.Title, &lesson.Content)
		if err != nil {
			return lessons, err
		}
		lessons = append(lessons, lesson)
	}

	return lessons, nil
}
