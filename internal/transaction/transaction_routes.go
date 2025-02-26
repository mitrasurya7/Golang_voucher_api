package transaction

import "github.com/gin-gonic/gin"

func TransactionRoutes(router *gin.RouterGroup) {
	transactionGroup := router.Group("/transaction/redemption")
	transactionGroup.POST("/", CreateTransactionHandler)
	transactionGroup.GET("/", GetTransactionByIDHandler)
}
