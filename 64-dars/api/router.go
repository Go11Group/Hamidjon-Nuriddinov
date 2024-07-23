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
		req := model.Company{}
		err := c.ShouldBindJSON(&req)
		if err != nil {
			log.Fatal(err)
		}
		resp := redisDB.Getredis(c, req.Company)
		c.JSON(http.StatusOK, resp)
	})

	return router
}
