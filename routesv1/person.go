package routesv1

import (
	"io/ioutil"

	"github.com/bartalcorn/apistarter/person"
	"github.com/gin-gonic/gin"
)

// Person is a sample route for POSTing data
func Person(c *gin.Context) {
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	if err != nil {
		c.JSON(400, gin.H{
			"status":   400,
			"protocol": "POST",
			"path":     "/",
			"message":  "error reading body",
			"error":    err.Error(),
		})
		return
	}

	// convert body
	valid, errs, person := person.ValidatePerson(value)
	if !valid {
		c.JSON(402, gin.H{
			"status":   402,
			"protocol": "POST",
			"path":     "/",
			"message":  "error parsing body",
			"error":    errs,
		})
		return
	}

	// return success
	c.JSON(200, gin.H{
		"status":   200,
		"protocol": "POST",
		"path":     "/",
		"message":  "person post route",
		"value":    person,
	})
}
