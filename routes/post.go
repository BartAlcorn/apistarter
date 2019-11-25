package routes

import (
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

// PostHomePage displays the default page via POST
func PostHomePage(c *gin.Context) {
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	if err != nil {
		c.JSON(500, gin.H{
			"protocol": "POST",
			"path":     "/",
			"message":  "error reading body",
			"error":    err.Error(),
		})
	}
	c.JSON(200, gin.H{
		"protocol": "POST",
		"path":     "/",
		"message":  "default route via post",
		"value":    string(value),
	})
}
