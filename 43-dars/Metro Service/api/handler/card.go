package handler

import (
	"mymode/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateCard(c *gin.Context) {
	card := model.CreateCard{}
	err := c.ShouldBindJSON(&card)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	err = h.Card.CreateCard(card)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusCreated, "Your information has been added successfully")
}

func (h *Handler) GetByIdCard(c *gin.Context) {
	id := c.Param("id")
	card, err := h.Card.GetByIdCard(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, card)
}

func (h *Handler) GetAllCards(c *gin.Context) {
	cards, err := h.Card.GetAllCards()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, cards)
}

func (h *Handler) UpdateCard(c *gin.Context) {
	id := c.Param("id")
	card := model.CreateCard{}
	err := c.ShouldBindJSON(&card)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	err = h.Card.UpdateCard(id, card)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, "Your information has been updated successfully")
}

func (h *Handler) DeleteCard(c *gin.Context) {
	id := c.Param("id")
	err := h.Card.DeleteCard(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, "Your data has been deleted successfully")
}



