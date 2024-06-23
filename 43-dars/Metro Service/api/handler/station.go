package handler

import (
	"mymode/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateStation(c *gin.Context) {
	station := model.CreateStation{}
	err := c.ShouldBindJSON(&station)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	err = h.Station.CreateStation(station)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusCreated, "Your information has been added successfully")
}

func (h *Handler) GetByIdStation(c *gin.Context) {
	id := c.Param("id")
	station, err := h.Station.GetByIdStation(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, station)
}

func (h *Handler) GetAllStations(c *gin.Context) {
	stations, err := h.Station.GetAllStations()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, stations)
}

func (h *Handler) UpdateStation(c *gin.Context) {
	id := c.Param("id")
	station := model.CreateStation{}
	err := c.ShouldBindJSON(&station)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	err = h.Station.UpdateStation(id, station)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, "Your information has been updated successfully")
}

func (h *Handler) DeleteStation(c *gin.Context) {
	id := c.Param("id")
	err := h.Station.DeleteStation(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, "Your data has been deleted successfully")
}
