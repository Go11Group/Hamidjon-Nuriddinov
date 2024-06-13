package api

import (
	"database/sql"
	"mymode/storage"

	"github.com/gin-gonic/gin"
)

type Handle struct {
	u *storage.NewUserRepo
	c *storage.NewCourseRepo
	l *storage.NewLessonRepo
	e *storage.NewEnrollmentRepo
}

func Router(db *sql.DB) *gin.Engine {
	router := gin.Default()

	var handle Handle
	var u ReqUser
	var c ReqCourse
	var l ReqLesson
	var e ReqEnrollment

	handle.u = storage.NewUser(db)
	handle.c = storage.NewCourse(db)
	handle.l = storage.NewLesson(db)
	handle.e = storage.NewEnrollment(db)

	u.ConnectorDb(handle)
	c.ConnectorDb(handle)
	l.ConnectorDb(handle)
	e.ConnectorDb(handle)

	user := router.Group("/user")
	user.POST("/createUser", u.CreateUser)
	user.GET("/readUser/:id", u.ReadUser)
	user.PUT("/updateUser/:id", u.UpdateUser)
	user.DELETE("/deleteUser/:id", u.DeleteUser)
	user.GET("/getAllUsers/:limit/:offset", u.GetAllUsers)

	course := router.Group("/course")
	course.POST("/createCourse", c.CreateCourse)
	course.GET("/readCourse/:id", c.ReadCourse)
	course.PUT("/updateCourse/:id", c.UpdateCourse)
	course.DELETE("/deleteCourse/:id", c.DeleteCourse)
	course.GET("/getAllCourses/:limit/:offset", c.GetAllCourses)

	lesson := router.Group("/lesson")
	lesson.POST("/createLesson", l.CreateLesson)
	lesson.GET("/readLesson/:id", l.ReadLesson)
	lesson.PUT("/updateLesson/:id", l.UpdateLesson)
	lesson.DELETE("/deleteLesson/:id", l.DeleteLesson)
	lesson.GET("/getAllLessons/:limit/:offset", l.GetAllLessons)

	enrollment:= router.Group("/enrollment")
	enrollment.POST("/createEnrollment", e.CreateEnrollment)
	enrollment.GET("/readEnrollment/:id", e.ReadEnrollment)
	enrollment.PUT("/updateEnrollment/:id", e.UpdateEnrollment)
	enrollment.DELETE("/deleteEnrollment/:id", e.DeleteEnrollment)
	enrollment.GET("/getAllEnrollments/:limit/:offset", e.GetAllEnrollments)

	return router
}
