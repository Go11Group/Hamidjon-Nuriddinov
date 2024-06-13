package storage

import (
	"database/sql"
	"mymode/module"
	"time"
)

type NewEnrollmentRepo struct {
	Db *sql.DB
}

func NewEnrollment(db *sql.DB) *NewEnrollmentRepo {
	return &NewEnrollmentRepo{Db: db}
}

func (E *NewEnrollmentRepo) CreateEnrollment(enrollment module.Enrollment) error {
	_, err := E.Db.Exec(`INSERT INTO 
							Enrollments(user_id, course_id) 
						VALUES($1, $2)`, enrollment.UserId, enrollment.CourseId)
	return err
}

func (E *NewEnrollmentRepo) ReadEnrollment(id string) (module.Enrollment, error) {
	enrollment := module.Enrollment{}
	err := E.Db.QueryRow(`SELECT enrollment_id, user_id, course_id, created_at FROM Enrollments
							WHERE enrollment_id = $1`, id).Scan(
		&enrollment.EnrollmentId, &enrollment.UserId, &enrollment.CourseId, &enrollment.CreatedAt)
	return enrollment, err
}

func (E *NewEnrollmentRepo) UpdateEnrollment(enrollment module.Enrollment, id string) error {
	_, err := E.Db.Exec(`UPDATE Enrollments SET 
							user_id, course_id
						WHERE 
							deleted_at is null and enrollment_id = $3`,
		enrollment.UserId, enrollment.CourseId, id)
	return err
}

func (E *NewEnrollmentRepo) DeleteEnrollment(id string) error {
	_, err := E.Db.Exec(`UPDATE Enrollments SET 
							deleted_at = $1
						WHERE 
							deleted_at is null and enrollment_id = $2`, time.Now(), id)
	return err
}

func (E *NewEnrollmentRepo) GetAllEnrollments(limit int, offset int) ([]module.Enrollment, error) {
	enrollments := []module.Enrollment{}
	rows, err := E.Db.Query(`SELECT 
								enrollment_id, user_id, course_id, created_at 
							FROM Enrollments
							LIMIT $1 
							OFFSET $2`, limit, offset)
	if err != nil {
		return enrollments, err
	}

	for rows.Next() {
		enrollment := module.Enrollment{}
		err = rows.Scan(&enrollment.EnrollmentId, &enrollment.UserId, &enrollment.CourseId, &enrollment.CreatedAt)
		if err != nil {
			return enrollments, err
		}
		enrollments = append(enrollments, enrollment)
	}

	return enrollments, nil
}
