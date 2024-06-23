package handler

import (
	"mymode/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateTerminal(c *gin.Context) {
	terminal := model.CreateTerminal{}
	err := c.ShouldBindJSON(&terminal)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	err = h.Terminal.CreateTerminal(terminal)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusCreated, "Your information has been added successfully")
}

func (h *Handler) GetByIdTerminal(c *gin.Context) {
	id := c.Param("id")
	terminal, err := h.Terminal.GetByIdTerminal(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, terminal)
}

func (h *Handler) GetAllTerminals(c *gin.Context) {
	terminals, err := h.Terminal.GetAllTerminals()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, terminals)
}

func (h *Handler) UpdateTerminal(c *gin.Context) {
	id := c.Param("id")
	terminal := model.CreateTerminal{}
	err := c.ShouldBindJSON(&terminal)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	err = h.Terminal.UpdateTerminal(id, terminal)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, "Your information has been updated successfully")
}

func (h *Handler) DeleteTerminal(c *gin.Context) {
	id := c.Param("id")
	err := h.Terminal.DeleteTerminal(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, "Your data has been deleted successfully")
}
