package middleware

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/bartalcorn/apistarter/responses"
)

// TokenAuth is a psudeo middleware function demonstrating token authentication.
func TokenAuth(c *gin.Context) {

	requiredToken := os.Getenv("API_TOKEN")

	// We want to make sure the token is set, bail if not
	if requiredToken == "" {
		log.Fatal("Please set API_TOKEN environment variable")
	}

	token := c.Request.FormValue("api_token")

	if token == "" {
		responses.RespondWithError(c, 401, "API token required")
		return
	}

	if token != requiredToken {
		fmt.Println("bad token", token)
		responses.RespondWithError(c, 403, "Authorization Failure")
		return
	}

	// Pass on to the next-in-chain
	c.Next()
}
