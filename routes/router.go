package routes

import (
	"github.com/gin-gonic/gin"
)

// Router is for root level routes
func Router(router *gin.Engine) *gin.Engine {
	// root level routes
	router.GET("/", Default)
	router.GET("/path/:name/:age", PathStrings)
	router.GET("/query", QueryStrings)
	router.PATCH("/", Post)
	router.POST("/", Post)
	router.PUT("/", Post)
	router.DELETE("/:name", PathStrings)

	return router
}
