package voucher

import (
	"golang-voucher-api/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateVoucherHandler(c *gin.Context) {
	var voucher Voucher
	if err := c.ShouldBindJSON(&voucher); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	createdVoucher, err := CreateNewVoucher(voucher.BrandID, voucher.Name, voucher.CostInPoint)
	if err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, "Failed to create voucher", nil)
		return
	}

	utils.JSONResponse(c, http.StatusCreated, "Voucher created successfully", createdVoucher)
}

func GetVoucherByIDHandler(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, "Invalid id", nil)
		return
	}

	voucher, err := GetSingleVoucherByID(uint(id))
	if err != nil {
		utils.JSONResponse(c, http.StatusNotFound, "Voucher not found", nil)
		return
	}

	utils.JSONResponse(c, http.StatusOK, "Voucher found", voucher)
}

func GetAllVoucherByBrandIDHandler(c *gin.Context) {
	brandIDStr := c.Query("brand_id")
	brandID, err := strconv.ParseUint(brandIDStr, 10, 64)
	if err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, "Invalid brand_id", nil)
		return
	}

	vouchers, err := GetAllVoucherByBrandID(uint(brandID))
	if err != nil {
		utils.JSONResponse(c, http.StatusNotFound, "Voucher not found", nil)
		return
	}

	utils.JSONResponse(c, http.StatusOK, "Voucher found", vouchers)
}
