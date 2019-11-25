package main

import (
	"github.com/gin-gonic/gin"

	"github.com/bartalcorn/apistarter/routes"
)

func main() {

	router := gin.Default()
	router.GET("/", routes.DefaultPage)
	router.GET("/query", routes.QueryStrings)
	router.GET("/path/:name/:age", routes.PathStrings)
	router.POST("/", routes.PostHomePage)
	router.PUT("/", routes.PostHomePage)
	router.PATCH("/", routes.PostHomePage)
	router.DELETE("/:name", routes.PathStrings)

	router.Run() // localhost:8080

}
