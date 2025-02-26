package brand_test

import (
	"golang-voucher-api/internal/brand"
	"golang-voucher-api/pkg/database"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
)

func TestCreateBrandHandler(t *testing.T) {

	database.DatabaseConnectTest()
	router := gin.Default()
	router.POST("/brand", brand.CreateBrandHandler)

	reqBody := `{"name": "Test Brand"}`

	req, _ := http.NewRequest("POST", "/brand", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}
