package voucher

import "github.com/gin-gonic/gin"

func VoucherRoutes(router *gin.RouterGroup) {
	voucherGroup := router.Group("/voucher")
	voucherGroup.POST("/", CreateVoucherHandler)
	voucherGroup.GET("/", GetVoucherByIDHandler)
	voucherGroup.GET("/brand", GetAllVoucherByBrandIDHandler)
}
