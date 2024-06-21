package api

import (
	"mymode/api/handler"

	"github.com/gin-gonic/gin"
)



func Router()*gin.Engine{
	router := gin.Default()

	user := router.Group("/users")
	user.GET("/readUser/:id", handler.ReadUser)

	return router
}