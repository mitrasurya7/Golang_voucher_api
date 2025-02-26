package routes

import (
	"golang-voucher-api/internal/brand"
	"golang-voucher-api/internal/transaction"
	"golang-voucher-api/internal/voucher"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api")
	brand.BrandRoutes(api)
	voucher.VoucherRoutes(api)
	transaction.TransactionRoutes(api)
}
