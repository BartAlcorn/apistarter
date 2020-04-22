package routes

import (
	"github.com/bartalcorn/apistarter/responses"
	"github.com/gin-gonic/gin"
)

// PathStrings displays the results of path string params
func PathStrings(c *gin.Context) {
	name := c.Param("name")
	age := c.Param("age")
	responses.RespondJSON(c, 200, gin.H{
		"name": name,
		"age":  age,
	})
}
