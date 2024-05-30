package student

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Student struct {
	Id     int
	Name   string
	Age    int
	Course string
}

type StudentRepo struct {
	Db *sql.DB
}

func NewStudentRepo(db *sql.DB) *StudentRepo {
	return &StudentRepo{Db: db}
}

func (S *StudentRepo) GetAllStudents() ([]Student, error) {
	var students = []Student{}
	rows, err := S.Db.Query(`Select S.student_id, S.name, S.age, C.name from Students as S 
	left join Student_course as Sc
	on S.student_id = Sc.student_id
	join Courses as C
	on C.course_id = Sc.course_id`)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var student Student
		err = rows.Scan(&student.Id, &student.Name, &student.Age, &student.Course)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}
	return students, nil
}

func (S *StudentRepo) GetByID(id int) (Student, error) {
	var student Student
	err := S.Db.QueryRow(`Select S.student_id, S.name, S.age, C.name from Students as S 
	left join Student_course as Sc
	on S.student_id = Sc.student_id
	join Courses as C
	on C.course_id = Sc.course_id
	where S.student_id = $1`, id).Scan(&student.Id, &student.Name, &student.Age, &student.Course)
	if err != nil {
		return Student{}, err
	}
	return student, nil
}


func (S *StudentRepo) Insert()error{
	var student Student
	fmt.Print("Id = ")
	fmt.Scan(&student.Id)
	fmt.Print("Name = ")
	fmt.Scan(&student.Name)
	fmt.Print("Age = ")
	fmt.Scan(&student.Age)
	_, err := S.Db.Exec(`INSERT INTO Students values ($1, $2, $3)`, student.Id, student.Name, student.Age)
	if err != nil{
		return err
	}
	fmt.Println("Ma'lumot muvaffaqiyatli qo'shildi.")
	return nil
}

func (S *StudentRepo) Update()error{
	var student Student
	fmt.Println("Qaysi userning ma'lumotlarini o'zgartirasiz?")
	fmt.Print("Id = ")
	fmt.Scan(&student.Id)
	fmt.Println("Yangi ma'lumotlarni kiriting.")
	fmt.Print("Name = ")
	fmt.Scan(&student.Name)
	fmt.Print("Age = ")
	fmt.Scan(&student.Age)
	_, err := S.Db.Exec(`Update Students Set name = $1, age = $2 where student_id = $3`, student.Name, student.Age, student.Id)
	if err != nil{
		panic(err)
	}
	fmt.Println("Ma'lumotlaringiz muvaffaqiyatli yangilandi.")
	return nil
}


func(S *StudentRepo) Delete(id int)error{
	_, err := S.Db.Exec(`Delete from Students where student_id = $1`, id)
	if err != nil{
		return err
	}
	fmt.Println("Ma'lumot muvaffaqiyatli o'chirildi.")
	return nil
}



