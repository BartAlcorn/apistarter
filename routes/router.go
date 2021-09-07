package routes

import (
	"github.com/gofiber/fiber/v2"
)

// Router is for root level routes
func Router(router fiber.Router) {
	// root level routes
	router.Get("/", Default())
	router.Get("/path/:name/:age", PathStrings())
	router.Get("/query", QueryStrings())
	router.Patch("/", Post())
	router.Post("/", Post())
	router.Put("/", Post())
	router.Delete("/:name", PathStrings())

}
