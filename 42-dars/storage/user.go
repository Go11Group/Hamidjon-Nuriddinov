package storage

import (
	"database/sql"
	"log"
	"mymode/model"
	"time"
)

type NewUserRepo struct {
	Db *sql.DB
}

func NewUser(db *sql.DB) *NewUserRepo {
	return &NewUserRepo{Db: db}
}

func (U *NewUserRepo) CreateUser(user model.User) error {
	_, err := U.Db.Exec(`INSERT INTO 
						Users(name, email, password) 
						VALUES($1, $2, $3)`, user.Name, user.Email, user.Password)
	return err
}

func (U *NewUserRepo) ReadUser(id string) (model.User, error) {
	user := model.User{}
	err := U.Db.QueryRow(`SELECT 
							  name, email, birthday FROM Users 
						  WHERE 
							  user_id = $1`, id).Scan(&user.Name, &user.Email, &user.Birthday)
	return user, err
}

func (U *NewUserRepo) UpdateUser(user model.User, id string) error {
	_, err := U.Db.Exec(`UPDATE Users SET 
							name = $1, email = $2, birthday = $3,  password = $4
						WHERE 
							deleted_at is null and user_id = $5`,
		user.Name, user.Email, user.Birthday, user.Password, id)
	return err
}

func (U *NewUserRepo) DeleteUser(id string) error {
	_, err := U.Db.Exec(`UPDATE Users SET 
							deleted_at = $1
						WHERE 
							deleted_at is null and user_id = $2`, time.Now(), id)
	return err
}

func (U *NewUserRepo) GetAllUsers(filter string, arr []interface{}) ([]model.User, error) {
	users := []model.User{}
	rows, err := U.Db.Query(`SELECT 
								user_id, name, email, birthday
							FROM 
								Users
							WHERE
								deleted_at is null` + filter, arr...)
	if err != nil {
		return users, err
	}

	for rows.Next() {
		user := model.User{}
		err = rows.Scan(&user.UserId, &user.Name, &user.Email, &user.Birthday)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (U *NewUserRepo) GetCoursesByUser(id string) (model.UserCourses, error) {
	courses := []model.ApiCourses{}
	rows, err := U.Db.Query(`SELECT
								C.course_id, C.title, C.description
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
								U.user_id = $1`, id)
	if err != nil {
		log.Fatal(err)
		return model.UserCourses{}, err
	}
	for rows.Next() {
		course := model.ApiCourses{}
		err = rows.Scan(&course.CourseId, &course.Title, &course.Description)
		if err != nil {
			log.Fatal(err)
			return model.UserCourses{}, err
		}
		courses = append(courses, course)
	}
	return model.UserCourses{
		UserId:  id,
		Courses: courses,
	}, err
}


func (U *NewUserRepo) SearchUsers(filter string, arr []interface{})([]model.UserSearch, error){
	users := []model.UserSearch{}
	rows, err := U.Db.Query(`SELECT
								user_id, name, email
							FROM 
								Users
							WHERE
								deleted_at is null` + filter, arr...)
	if err != nil {
		log.Fatal(err)
		return users, err
	}
	for rows.Next() {
		user := model.UserSearch{}
		err = rows.Scan(&user.UserId, &user.Name, &user.Email)
		if err != nil {
			log.Fatal(err)
			return users, err
		}
		users = append(users, user)
	}
	return users, err
}
