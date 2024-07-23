package api

import (
	redisDB "homework/Redis"
	"homework/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()

	router.GET("/getPrice", func(c *gin.Context) {
		req := model.CompanyName{}
		err := c.ShouldBindJSON(&req)
		if err != nil {
			log.Fatal(err)
		}
		resp := redisDB.Getredis(c, req.Name)
		c.JSON(http.StatusOK, resp)
	})

	return router
}
