package api

import (
	"mymode/api/handler"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func Router(conn grpc.ClientConn) *gin.Engine {
	router := gin.Default()
	h := handler.NewHandlerRepo(conn)

	weather := router.Group("/weather")
	weather.GET("/getCurrentWeather/:time", h.GetCurrentWeather)
	weather.GET("/getWeatherForecast/", h.GetWeatherForecast)
	weather.POST("reportWeatherCondition", h.ReportWeatherCondition)

	transport := router.Group("/transport")
	transport.GET("/getBusSchedule/:number", h.GetBusSchedule)
	transport.GET("/trackBusLocation/:number", h.TrackBusLocation)
	transport.POST("/reportTrafficJam", h.ReportTrafficJam)

	return router
}
