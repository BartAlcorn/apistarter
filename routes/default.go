package routes

import (
	"github.com/gin-gonic/gin"
)

// DefaultPage displays the default page via GET
func DefaultPage(c *gin.Context) {
	c.JSON(200, gin.H{
		"protocol": "GET",
		"path":    "/",
		"message": "default home page",
	})
}
