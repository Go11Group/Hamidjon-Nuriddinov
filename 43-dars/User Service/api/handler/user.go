package handler

import (
	"encoding/json"
	"log"
	"mymode/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateUser(c *gin.Context) {
	user := model.CreateUser{}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		log.Fatal(err)
	}
	err = h.User.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		log.Fatal(err)
	}
	c.JSON(http.StatusCreated, "Your information has been added successfully")
}

func (h *Handler) GetByIdUser(c *gin.Context) {
	id := c.Param("id")
	user, err := h.User.GetByIdUser(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, user)
}

func (h *Handler) GetAllUsers(c *gin.Context) {
	users, err := h.User.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, users)
}

func (h *Handler) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	user := model.CreateUser{}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		log.Fatal(err)
	}
	err = h.User.UpdateUser(id, user)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, "Your information has been updated successfully")
}

func (h *Handler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	err := h.User.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, "Your data has been deleted successfully")
}

func (h *Handler) GetBalance(c *gin.Context) {
	userId := c.Param("userId")
	cardId := c.Param("cardId")
	user, err := h.User.GetByIdUser(userId)
	emptyUser := model.CreateUser{}
	if err != nil || user == emptyUser {
		c.JSON(http.StatusBadRequest, "The user does not exist in the database")
		return
	}
	resp, err := http.Get("http://localhost:8082/card/getBalance/" + userId + "/" + cardId)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	balance := model.UserBalance{}
	err = json.NewDecoder(resp.Body).Decode(&balance)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, balance)
}
