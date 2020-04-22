package main

import (
	"fmt"
	"time"

	"github.com/bartalcorn/apistarter/middleware"
	"github.com/bartalcorn/apistarter/routes"
	"github.com/bartalcorn/apistarter/routesv1"
	"github.com/gin-gonic/gin"
)

func main() {

	// gin.DisableConsoleColor()

	router := gin.Default()
	// router.Use(gin.Logger())
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	router.Use(middleware.TokenAuth)

	// routes
	router.Use(gin.Recovery())
	router = routes.Router(router)
	routesv1.Router(router)

	// listener
	err := router.Run() // localhost:8080
	if err != nil {
		fmt.Print("Oops", err)
	}
}
