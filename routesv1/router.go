package routesv1

import (
	"github.com/bartalcorn/apistarter/routes"
	"github.com/gofiber/fiber/v2"
)

func DefaultRouter(app fiber.Router) {
	app.Get("/", routes.Default())
	app.Get("/path/:name/:age", routes.PathStrings())
	app.Get("/query", routes.QueryStrings())
	app.Post("/person", Person())
}
