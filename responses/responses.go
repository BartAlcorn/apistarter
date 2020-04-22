package responses

import "github.com/gin-gonic/gin"

// RespondJSON is a convenience function for returning a JSON response.
func RespondJSON(c *gin.Context, code int, message interface{}) {
	c.JSON(code, gin.H{
		"status":   code,
		"protocol": c.Request.Method,
		"path":     c.Request.URL.Path,
		"success":  message,
	})
}

// RespondWithError is a convenience function for returning an error.
// Aborts further processing and returns immediately.
func RespondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{
		"status":   code,
		"protocol": c.Request.Method,
		"path":     c.Request.URL.Path,
		"error":    message,
	})
}
