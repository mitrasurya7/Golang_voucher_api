package transaction_test

import (
	"golang-voucher-api/internal/transaction"
	"golang-voucher-api/pkg/database"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/transaction/redemption", transaction.CreateTransactionHandler)
	router.GET("/transaction/redemption", transaction.GetTransactionByIDHandler)
	return router
}
func TestCreateTransaction(t *testing.T) {
	database.DatabaseConnectTest()
	router := setupRouter()
	reqBody := `{"voucher_id":1,"quantity":1,"total_point":10000}`

	req, _ := http.NewRequest("POST", "/transaction/redemption", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestGetTransactionByIDHandler(t *testing.T) {
	database.DatabaseConnectTest()
	router := setupRouter()

	req := httptest.NewRequest("GET", "/transaction/redemption?transactionId=2", nil)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
