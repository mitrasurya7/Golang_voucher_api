package utils

import (
	"github.com/gin-gonic/gin"
)

func JSONResponse(c *gin.Context, statusCode int, message string, data any) {
	c.JSON(statusCode, gin.H{
		"statusCode": statusCode,
		"message":    message,
		"data":       data,
	})
}
