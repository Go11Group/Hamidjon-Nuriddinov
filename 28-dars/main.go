package main

import (
	"mymod/packages"
	"mymod/packages/course"
	"mymod/packages/student"
)

func main() {
	db, err := packages.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// st := student.NewStudentRepo(db)
	// students, err := st.GetAllStudents()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(students)

	// student, err := st.GetByID(1)
	// if err != nil{
	// 	panic(err)
	// }
	// fmt.Println(student)

	// err = st.Insert()
	// if err != nil{
	// 	panic(err)
	// }

	// err = st.Update()
	// if err != nil{
	// 	panic(err)
	// }

	// err = st.Delete(45)
	// if err != nil{
	// 	panic(err)
	// }

	// c := course.NewCourseRepo(db)

	// courses, err := c.GetAllCourses()
	// if err != nil{
	// 	panic(err)
	// }
	// fmt.Println(courses)

	// course, err := c.GetByID(1)
	// if err != nil{
	// 	panic(err)
	// }
	// fmt.Println(course)

	// err = c.Insert()
	// if err != nil{
	// 	panic(err)
	// }

	// err = c.Update()
	// if err != nil{
	// 	panic(err)
	// }

	// err = c.Delete(15)
	// if err != nil{
	// 	panic(err)
	// }
}
