package transaction

import (
	"golang-voucher-api/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateTransactionHandler(c *gin.Context) {
	var transaction Transaction
	if err := c.ShouldBindJSON(&transaction); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	createdTransaction, err := CreateNewTransaction(transaction.VoucherID, transaction.Quantity, transaction.TotalPoint)
	if err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, "Failed to create transaction", nil)
		return
	}
	utils.JSONResponse(c, http.StatusCreated, "Transaction created successfully", createdTransaction)
}

func GetTransactionByIDHandler(c *gin.Context) {
	idStr := c.Query("transactionId")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, "Invalid transactionId", nil)
		return
	}

	transaction, err := GetDetailTransactionByID(uint(id))
	if err != nil {
		utils.JSONResponse(c, http.StatusNotFound, "Transaction not found", nil)
		return
	}

	utils.JSONResponse(c, http.StatusOK, "Transaction found", transaction)
}
