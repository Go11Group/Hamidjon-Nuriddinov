package api

import (
	"database/sql"
	"mymode/api/handler"

	"github.com/gin-gonic/gin"
)

func Router(db *sql.DB) *gin.Engine {
	router := gin.Default()
	h := handler.NewHandler(db)

	user := router.Group("/user")
	user.POST("/createUser", h.CreateUser)
	user.GET("/getByIdUser/:id", h.GetByIdUser)
	user.GET("/getAllUsers", h.GetAllUsers)
	user.PUT("/updateUser/:id", h.UpdateUser)
	user.DELETE("/deleteUser/:id", h.DeleteUser)

	return router
}
