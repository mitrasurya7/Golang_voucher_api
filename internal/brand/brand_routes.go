package brand

import "github.com/gin-gonic/gin"

func BrandRoutes(router *gin.RouterGroup) {
	brandGroup := router.Group("/brand")
	brandGroup.POST("/", CreateBrandHandler)
}
