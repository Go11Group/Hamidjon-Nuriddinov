package lessons

import (
	"database/sql"
	"time"
)

type NewLessonRepo struct {
	Db *sql.DB
}

func NewLesson(db *sql.DB) *NewLessonRepo {
	return &NewLessonRepo{Db: db}
}

func (L *NewLessonRepo) Create(lesson Lesson) error {
	_, err := L.Db.Exec(`INSERT INTO 
						Lessons(title, content) 
						VALUES($1, $2)`, lesson.Title, lesson.Content)
	return err
}

func (L *NewLessonRepo) Read(id string) (Lesson, error) {
	lesson := Lesson{}
	err := L.Db.QueryRow(`SELECT 
							  title, content, created_at FROM Lessons
						  WHERE 
							  lesson_id = $1`, id).Scan(&lesson.Title, &lesson.Content)
	return lesson, err
}

func (L *NewLessonRepo) Update(lesson Lesson, id string) error {
	_, err := L.Db.Exec(`UPDATE Lessons SET 
							title = $1, content = $2
						WHERE 
							deleted_at is null lesson_id = $3`,
		lesson.Title, lesson.Content, id)
	return err
}

func (L *NewLessonRepo) Delete(id string) error {
	_, err := L.Db.Exec(`UPDATE Lessons SET 
							deleted_at = $1
						WHERE 
							deleted_at is null and lesson_id = $2`, time.Now(), id)
	return err
}

func (L *NewLessonRepo) GetAll(limit int, offset int) ([]Lesson, error) {
	lessons := []Lesson{}
	rows, err := L.Db.Query(`SELECT 
								title, content, created_at 
							FROM Lessons
							LIMIT $1 
							OFFSET $2`, limit, offset)
	if err != nil {
		return lessons, err
	}

	for rows.Next() {
		lesson := Lesson{}
		err = rows.Scan(&lesson.Title, &lesson.Content, &lesson.CreatedAt)
		if err != nil {
			return lessons, err
		}
		lessons = append(lessons, lesson)
	}

	return lessons, nil
}
