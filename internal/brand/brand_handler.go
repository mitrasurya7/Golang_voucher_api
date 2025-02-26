package brand

import (
	"fmt"
	"golang-voucher-api/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateBrandHandler(c *gin.Context) {
	var brand Brand
	if err := c.ShouldBindJSON(&brand); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	createdBrand, err := CreateNewBrand(brand.Name)
	if err != nil {
		fmt.Println("Database Error:", err)
		utils.JSONResponse(c, http.StatusInternalServerError, "Failed to create brand", nil)
		return
	}

	utils.JSONResponse(c, http.StatusCreated, "Brand created successfully", createdBrand)
}
