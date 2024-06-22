package api

import (
	"mymode/api/handler"

	"github.com/gin-gonic/gin"
)



func Router()*gin.Engine{
	router := gin.Default()

	user := router.Group("/users")
	user.POST("/createUser", handler.CreateUser)
	user.GET("/readUser/:id", handler.ReadUser)
	user.GET("/getAllUsers", handler.GetAllUsers)
	user.PUT("/updateUser/:id", handler.UpdateUser)
	user.DELETE("deleteUser/:id", handler.DeleteUser)

	course := router.Group("/courses")
	course.POST("/createCourse", handler.CreateCourse)
	course.GET("/readCourse/:id", handler.ReadCourse)
	course.GET("/getAllCourses", handler.GetAllCourses)
	course.PUT("/updateCourse/:id", handler.UpdateCourse)
	course.DELETE("deleteCourse/:id", handler.DeleteCourse)

	lesson := router.Group("/lessons")
	lesson.POST("/createLesson", handler.CreateLesson)
	lesson.GET("/readLesson/:id", handler.ReadLesson)
	lesson.GET("/getAllLessons", handler.GetAllLessons)
	lesson.PUT("/updateLesson/:id", handler.UpdateLesson)
	lesson.DELETE("deleteLesson/:id", handler.DeleteLesson)

	enrollment := router.Group("/enrollments")
	enrollment.POST("/createEnrollment", handler.CreateEnrollment)
	enrollment.GET("/readEnrollment/:id", handler.ReadEnrollment)
	enrollment.GET("/getAllEnrollments", handler.GetAllEnrollments)
	enrollment.PUT("/updateEnrollment/:id", handler.UpdateEnrollment)
	enrollment.DELETE("deleteEnrollment/:id", handler.DeleteEnrollment)

	return router
}