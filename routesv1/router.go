package routesv1

import (
	"github.com/bartalcorn/apistarter/routes"
	"github.com/gin-gonic/gin"
)

func Router(router *gin.Engine) {
	// /v1/ level routes
	v1 := router.Group("/v1")
	{
		v1.GET("/", routes.Default)
		v1.GET("/path/:name/:age", routes.PathStrings)
		v1.GET("/query", routes.QueryStrings)
		v1.POST("/person", Person)
	}
}
