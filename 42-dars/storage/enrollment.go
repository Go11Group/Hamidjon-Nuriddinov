package storage

import (
	"database/sql"
	"mymode/model"
	"time"
)

type NewEnrollmentRepo struct {
	Db *sql.DB
}

func NewEnrollment(db *sql.DB) *NewEnrollmentRepo {
	return &NewEnrollmentRepo{Db: db}
}

func (E *NewEnrollmentRepo) CreateEnrollment(enrollment model.Enrollment) error {
	_, err := E.Db.Exec(`INSERT INTO 
							Enrollments(user_id, course_id) 
						VALUES($1, $2)`, enrollment.UserId, enrollment.CourseId)
	return err
}

func (E *NewEnrollmentRepo) ReadEnrollment(id string) (model.Enrollment, error) {
	enrollment := model.Enrollment{}
	err := E.Db.QueryRow(`SELECT enrollment_id, user_id, course_id FROM Enrollments
							WHERE enrollment_id = $1`, id).Scan(
		&enrollment.EnrollmentId, &enrollment.UserId, &enrollment.CourseId)
	return enrollment, err
}

func (E *NewEnrollmentRepo) UpdateEnrollment(enrollment model.Enrollment, id string) error {
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

func (E *NewEnrollmentRepo) GetAllEnrollments(filter string, arr []interface{}) ([]model.Enrollment, error) {
	enrollments := []model.Enrollment{}
	rows, err := E.Db.Query(`SELECT 
								enrollment_id, user_id, course_id, enrollment_date, created_at
							FROM 
								Enrollments
							WHERE
								deleted_at is null` + filter, arr...)
	if err != nil {
		return enrollments, err
	}

	for rows.Next() {
		enrollment := model.Enrollment{}
		err = rows.Scan(&enrollment.EnrollmentId, &enrollment.UserId, &enrollment.CourseId)
		if err != nil {
			return enrollments, err
		}
		enrollments = append(enrollments, enrollment)
	}

	return enrollments, nil
}


