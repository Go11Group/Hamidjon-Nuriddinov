package api

import (
	"mymode/api/handler"

	"github.com/gin-gonic/gin"
)


func Router()*gin.Engine{
	router := gin.Default()

	user := router.Group("/user")
	user.POST("/createUser", handler.CreateUser)
	user.GET("/getByIdUser/:id", handler.GetByIdUser)
	user.GET("/getAllUsers", handler.GetAllUsers)
	user.PUT("/updateUser/:id", handler.UpdateUser)
	user.DELETE("/deleteUser/:id", handler.DeleteUser)

	station := router.Group("/station")
	station.POST("/createStation", handler.CreateStation)
	station.GET("/getByIdStation/:id", handler.GetByIdStation)
	station.GET("/getAllStations", handler.GetAllStations)
	station.PUT("/updateStation/:id", handler.UpdateStation)
	station.DELETE("/deleteStation/:id", handler.DeleteStation)

	card := router.Group("/card")
	card.POST("/createCard", handler.CreateCard)
	card.GET("/getByIdCard/:id", handler.GetByIdCard)
	card.GET("/getAllCards", handler.GetAllCards)
	card.PUT("/updateCard/:id", handler.UpdateCard)
	card.DELETE("/deleteCard/:id", handler.DeleteCard)

	terminal := router.Group("/terminal")
	terminal.POST("/createTerminal", handler.CreateTerminal)
	terminal.GET("/getByIdTerminal/:id", handler.GetByIdTerminal)
	terminal.GET("/getAllTerminals", handler.GetAllTerminals)
	terminal.PUT("/updateTerminal/:id", handler.UpdateTerminal)
	terminal.DELETE("/deleteTerminal/:id", handler.DeleteTerminal)

	transaction := router.Group("/transaction")
	transaction.POST("/createTransaction", handler.CreateTransaction)
	transaction.GET("/getByIdTransaction/:id", handler.GetByIdTransaction)
	transaction.GET("/getAllTransactions", handler.GetAllTransactions)
	transaction.PUT("/updateTransaction/:id", handler.UpdateTransaction)
	transaction.DELETE("/deleteTransaction/:id", handler.DeleteTransaction)

	return router
}