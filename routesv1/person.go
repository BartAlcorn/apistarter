package routesv1

import (
	"io/ioutil"

	"github.com/bartalcorn/apistarter/person"
	"github.com/bartalcorn/apistarter/responses"
	"github.com/gin-gonic/gin"
)

// Person is a sample route for POSTing data
func Person(c *gin.Context) {
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	if err != nil {
		responses.RespondWithError(c, 400, gin.H{
			"message": "error reading body",
			"error":   err.Error(),
		})
		return
	}

	// convert body
	valid, errs, person := person.ValidatePerson(value)
	if !valid {
		responses.RespondWithError(c, 402, gin.H{
			"message": "error parsing body",
			"error":   errs,
		})
		return
	}

	// return success
	responses.RespondJSON(c, 200, gin.H{
		"message": "person post route",
		"person":  person,
	})

}
