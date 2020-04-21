package routes

import (
	"github.com/gin-gonic/gin"
)

// QueryStrings displays the results of query string params
func QueryStrings(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age")
	c.JSON(200, gin.H{
		"protocol": "GET",
		"path":    "/query",
		"message": "querystrings via get",
		"name":    name,
		"age":     age,
	})
}
