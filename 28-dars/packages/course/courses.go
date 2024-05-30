package course

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Course struct {
	Id   int
	Name string
}

type CourseRepo struct {
	Db *sql.DB
}

func NewCourseRepo(db *sql.DB) *CourseRepo {
	return &CourseRepo{Db: db}
}

func (S *CourseRepo) GetAllCourses() ([]Course, error) {
	var courses = []Course{}
	rows, err := S.Db.Query(`Select * from Courses`)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var course Course
		err = rows.Scan(&course.Id, &course.Name)
		if err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}
	return courses, nil
}

func (S *CourseRepo) GetByID(id int) (Course, error) {
	var course Course
	err := S.Db.QueryRow(`Select * from Courses
	where Course_id = $1`, id).Scan(&course.Id, &course.Name)
	if err != nil {
		return Course{}, err
	}
	return course, nil
}

func (S *CourseRepo) Insert() error {
	var course Course
	fmt.Print("Id = ")
	fmt.Scan(&course.Id)
	fmt.Print("Name = ")
	fmt.Scan(&course.Name)
	_, err := S.Db.Exec(`INSERT INTO Courses values ($1, $2)`, course.Id, course.Name)
	if err != nil {
		return err
	}
	fmt.Println("Ma'lumot muvaffaqiyatli qo'shildi.")
	return nil
}

func (S *CourseRepo) Update() error {
	var course Course
	fmt.Println("Qaysi kursning ma'lumotlarini o'zgartirasiz?")
	fmt.Print("Id = ")
	fmt.Scan(&course.Id)
	fmt.Println("Yangi ma'lumotlarni kiriting.")
	fmt.Print("Name = ")
	fmt.Scan(&course.Name)
	_, err := S.Db.Exec(`Update Courses Set name = $1 where course_id = $2`, course.Name, course.Id)
	if err != nil {
		panic(err)
	}
	fmt.Println("Ma'lumotlaringiz muvaffaqiyatli yangilandi.")
	return nil
}

func (S *CourseRepo) Delete(id int) error {
	_, err := S.Db.Exec(`Delete from Courses where Course_id = $1`, id)
	if err != nil {
		return err
	}
	fmt.Println("Ma'lumot muvaffaqiyatli o'chirildi.")
	return nil
}
