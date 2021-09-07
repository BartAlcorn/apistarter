package routes

import (
	"github.com/gofiber/fiber/v2"
)

// PathStrings displays the results of path string params
func PathStrings() fiber.Handler {
	return func(c *fiber.Ctx) error {
		name := c.Params("name")
		age := c.Params("age")
		return c.JSON(&fiber.Map{
			"name": name,
			"age":  age,
		})
	}
}
