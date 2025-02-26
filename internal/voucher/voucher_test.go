package voucher_test

import (
	"golang-voucher-api/internal/voucher"
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

	router.POST("/voucher", voucher.CreateVoucherHandler)
	router.GET("/voucher", voucher.GetVoucherByIDHandler)
	router.GET("/voucher/brand", voucher.GetAllVoucherByBrandIDHandler)

	return router
}

func TestCreateVoucherHandler(t *testing.T) {
	database.DatabaseConnectTest()
	router := setupRouter()

	reqBody := `{"brand_id":3,"name": "Voucher Test 10.000","cost_in_point": 30000}`
	req := httptest.NewRequest("POST", "/voucher", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestGetVoucherByIDHandler(t *testing.T) {
	database.DatabaseConnectTest()
	router := setupRouter()

	req := httptest.NewRequest("GET", "/voucher?id=1", nil)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetAllVoucherByBrandIDHandler(t *testing.T) {
	database.DatabaseConnectTest()
	router := setupRouter()

	req := httptest.NewRequest("GET", "/vouchers?brand_id=1", nil)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
