package courses

import (
	"database/sql"
	"time"
)

type NewCourseRepo struct {
	Db *sql.DB
}

func NewCourse(db *sql.DB) *NewCourseRepo {
	return &NewCourseRepo{Db: db}
}

func (C *NewCourseRepo) Create(course Course) error {
	_, err := C.Db.Exec(`INSERT INTO 
						Courses(title, description) 
						VALUES($1, $2)`, course.Title, course.Description)
	return err
}

func (C *NewCourseRepo) Read(id string) (Course, error) {
	course := Course{}
	err := C.Db.QueryRow(`SELECT 
							  title, description, created_at FROM Courses
						  WHERE 
							  course_id = $1`, id).Scan(&course.Title, &course.Description)
	return course, err
}

func (C *NewCourseRepo) Update(course Course, id string) error {
	_, err := C.Db.Exec(`UPDATE Courses SET 
							title = $1, description = $2
						WHERE 
							deleted_at is null course_id = $3`,
		course.Title, course.Description, id)
	return err
}

func (C *NewCourseRepo) Delete(id string) error {
	_, err := C.Db.Exec(`UPDATE Courses SET 
							deleted_at = $1
						WHERE 
							deleted_at is null and course_id = $2`, time.Now(), id)
	return err
}

func (C *NewCourseRepo) GetAll(limit int, offset int) ([]Course, error) {
	courses := []Course{}
	rows, err := C.Db.Query(`SELECT 
								title, description, created_at 
							FROM Courses
							LIMIT $1 
							OFFSET $2`, limit, offset)
	if err != nil {
		return courses, err
	}

	for rows.Next() {
		course := Course{}
		err = rows.Scan(&course.Title, &course.Description, &course.CreatedAt)
		if err != nil {
			return courses, err
		}
		courses = append(courses, course)
	}

	return courses, nil
}
