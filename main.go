package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/bartalcorn/apistarter/routesv1"
)

func main() {
	app := fiber.New()

	// Middleware
	app.Use(cors.New())
	app.Use(recover.New())

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Hello, World 👋!")
	// })

	routesv1.DefaultRouter(app)

	log.Fatal(app.Listen(":3000"))

}
