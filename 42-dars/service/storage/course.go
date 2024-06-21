package storage

import (
	"database/sql"
	"log"
	"mymode/model"
	"time"

)

type NewCourseRepo struct {
	Db *sql.DB
}

func NewCourse(db *sql.DB) *NewCourseRepo {
	return &NewCourseRepo{Db: db}
}

func (C *NewCourseRepo) CreateCourse(course model.Course) error {
	_, err := C.Db.Exec(`INSERT INTO 
						Courses(title, description) 
						VALUES($1, $2)`, course.Title, course.Description)
	return err
}

func (C *NewCourseRepo) ReadCourse(id string) (model.Course, error) {
	course := model.Course{}
	err := C.Db.QueryRow(`SELECT 
							  title, description
						  FROM 
						  	  Courses
						  WHERE 
							  course_id = $1`, id).Scan(&course.Title, &course.Description)
	return course, err
}

func (C *NewCourseRepo) UpdateCourse(course model.Course, id string) error {
	_, err := C.Db.Exec(`UPDATE Courses SET 
							title = $1, description = $2
						WHERE 
							deleted_at is null and course_id = $3`,
		course.Title, course.Description, id)
	return err
}

func (C *NewCourseRepo) DeleteCourse(id string) error {
	_, err := C.Db.Exec(`UPDATE Courses SET 
							deleted_at = $1
						WHERE 
							deleted_at is null and course_id = $2`, time.Now(), id)
	return err
}

func (C *NewCourseRepo) GetAllCourses(filter string, arr []interface{}) ([]model.Course, error) {
	courses := []model.Course{}
	rows, err := C.Db.Query(`SELECT 
								course_id, title, description, created_at
							FROM 
								Lessons
							WHERE
								deleted_at is null` + filter, arr...)
	if err != nil {
		return courses, err
	}

	for rows.Next() {
		course := model.Course{}
		err = rows.Scan(&course.Title, &course.Description)
		if err != nil {
			return courses, err
		}
		courses = append(courses, course)
	}

	return courses, nil
}

func (C *NewCourseRepo) GetLessonsByCourse(id string) (model.CourseLesons, error) {
	lessons := []model.ApiLesson{}
	rows, err := C.Db.Query(`SELECT
								L.lesson_id, L.title, L.content
							FROM 
								Course as C
							JOIN 
								Lesson as L
							ON 
								C.cours_id = L.course_id
							WHERE
								C.course_id = $1`, id)
	if err != nil {
		log.Fatal(err)
		return  model.CourseLesons{}, err
	}
	for rows.Next() {
		var lesson model.ApiLesson 
		err = rows.Scan(&lesson.LessonId, &lesson.Title, &lesson.Content)
		if err != nil {
			log.Fatal(err)
			return model.CourseLesons{}, err
		}
		lessons = append(lessons, lesson)
	}

	return model.CourseLesons{
		CourseId: id,
		Lessons: lessons,
	}, err
}

func (C *NewCourseRepo) GetEnrolledUsersByCourse(course_id string)(model.CourseUsers, error){
	users := []model.ApiUsers{}
	rows, err := C.Db.Query(`SELECT 
								U.user_id, U.name, U.email
							FROM
								Users as U
							JOIN
								Enrollments as E
							ON
								U.user_id = E.user_id
							JOIN 
								Courses as C
							ON
								C.course_id = E.course_id
							WHERE
								C.course_id = $1`, course_id)
	if err != nil{
		return model.CourseUsers{}, err
	}	
	for rows.Next() {
		var user = model.ApiUsers{}
		err := rows.Scan(&user.UserId, &user.Name, &user.Email)
		if err != nil{
			return model.CourseUsers{}, err
		}
		users = append(users, user)
	}	
	return model.CourseUsers{
		CourseId: course_id,
		EnrolledUser: users,
	}, err	
}
