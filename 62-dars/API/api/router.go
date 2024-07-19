package api

import (
	"api/api/handler"
	"api/api/middleware"

	"github.com/gin-gonic/gin"
)


func Router()*gin.Engine{
	router := gin.Default()
	h := handler.NewHandler()

	router.GET("/getCenter/:id", h.GetCenterById)

	admin := router.Group("/")
	admin.Use(middleware.AdminMiddleware())
	admin.POST("/createCenter", h.CreateCenter)
	admin.PUT("/updateCenter", h.UpdateCenter)
	admin.DELETE("/deleteCenter/:id", h.DeleteCenter)

	return router
}