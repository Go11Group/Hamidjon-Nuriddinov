package api

import (
	"database/sql"
	"mymode/api/handler"

	"github.com/gin-gonic/gin"
)


func Router(db *sql.DB)*gin.Engine{
	h := handler.NewHandler(db)

	router := gin.Default()

	card := router.Group("/card")
	card.POST("/createCard", h.CreateCard)
	card.GET("/getByIdCard/:id", h.GetByIdCard)
	card.GET("/getAllCards", h.GetAllCards)
	card.PUT("/updateCard/:id", h.UpdateCard)
	card.DELETE("/deleteCard/:id", h.DeleteCard)
	card.GET("/getBalance/:userId/:cardId", h.GetBalance)

	station := router.Group("/station")
	station.POST("/createStation", h.CreateStation)
	station.GET("/getByIdStation/:id", h.GetByIdStation)
	station.GET("/getAllStations", h.GetAllStations)
	station.PUT("/updateStation/:id", h.UpdateStation)
	station.DELETE("/deleteStation/:id", h.DeleteStation)

	terminal := router.Group("/terminal")
	terminal.POST("/createTerminal", h.CreateTerminal)
	terminal.GET("/getByIdTerminal/:id", h.GetByIdTerminal)
	terminal.GET("/getAllTerminals", h.GetAllTerminals)
	terminal.PUT("/updateTerminal/:id", h.UpdateTerminal)
	terminal.DELETE("/deleteTerminal/:id", h.DeleteTerminal)

	transaction := router.Group("/transaction")
	transaction.POST("/createTransaction", h.CreateTransaction)
	transaction.GET("/getByIdTransaction/:id", h.GetByIdTransaction)
	transaction.GET("/getAllTransactions", h.GetAllTransactions)
	transaction.PUT("/updateTransaction/:id", h.UpdateTransaction)
	transaction.DELETE("/deleteTransaction/:id", h.DeleteTransaction)
	return router
}