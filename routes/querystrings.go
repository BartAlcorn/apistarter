package routes

import (
	"github.com/bartalcorn/apistarter/responses"
	"github.com/gin-gonic/gin"
)

// QueryStrings displays the results of query string params
func QueryStrings(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age")
	responses.RespondJSON(c, 200, gin.H{
		"name": name,
		"age":  age,
	})
}
