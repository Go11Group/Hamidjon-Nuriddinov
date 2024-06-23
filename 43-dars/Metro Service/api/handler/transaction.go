package handler

import (
	"mymode/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateTransaction(c *gin.Context) {
	transaction := model.CreateTransaction{}
	err := c.ShouldBindJSON(&transaction)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	err = h.Transaction.CreateTransaction(transaction)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusCreated, "Your information has been added successfully")
}

func (h *Handler) GetByIdTransaction(c *gin.Context) {
	id := c.Param("id")
	transaction, err := h.Transaction.GetByIdTransaction(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, transaction)
}

func (h *Handler) GetAllTransactions(c *gin.Context) {
	transactions, err := h.Transaction.GetAllTransactions()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, transactions)
}

func (h *Handler) UpdateTransaction(c *gin.Context) {
	id := c.Param("id")
	transaction := model.CreateTransaction{}
	err := c.ShouldBindJSON(&transaction)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	err = h.Transaction.UpdateTransaction(id, transaction)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, "Your information has been updated successfully")
}

func (h *Handler) DeleteTransaction(c *gin.Context) {
	id := c.Param("id")
	err := h.Transaction.DeleteTransaction(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, "Your data has been deleted successfully")
}
