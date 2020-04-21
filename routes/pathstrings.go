package routes

import (
	"github.com/gin-gonic/gin"
)

// PathStrings displays the results of path string params
func PathStrings(c *gin.Context) {
	name := c.Param("name")
	age := c.Param("age")
	c.JSON(200, gin.H{
		"protocol": "GET",
		"path":    "/path/:name/:age",
		"message": "path strings via get",
		"name":    name,
		"age":     age,
	})
}
