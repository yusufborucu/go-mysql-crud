package utils

import "github.com/gin-gonic/gin"

func JSONResponse(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{"message": message})
}
