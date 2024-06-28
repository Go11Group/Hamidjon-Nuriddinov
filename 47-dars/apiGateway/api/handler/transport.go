package handler

import (
	"context"
	t "mymode/protos/transport/protos"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *handler) GetBusSchedule(c *gin.Context) {
	number, err := strconv.Atoi(c.Param("number"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	transport := t.NewTransportClient(&h.Connection)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	resp, err := transport.GetBusSchedule(ctx, &t.Number{Number: int32(number)})
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *handler) TrackBusLocation(c *gin.Context) {
	number, err := strconv.Atoi(c.Param("number"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	transport := t.NewTransportClient(&h.Connection)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	resp, err := transport.TrackBusLocation(ctx, &t.Number{Number: int32(number)})
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *handler) ReportTrafficJam(c *gin.Context) {
	station := c.Query("station")
	transport := t.NewTransportClient(&h.Connection)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	resp, err := transport.ReportTrafficJam(ctx, &t.Location{Station: station})
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, resp.Status)
}
