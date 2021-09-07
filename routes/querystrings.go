package routes

import (
	"github.com/gofiber/fiber/v2"
)

// QueryStrings displays the results of query string params
func QueryStrings() fiber.Handler {
	return func(c *fiber.Ctx) error {
		name := c.Query("name")
		age := c.Query("age")
		_ = c.JSON(&fiber.Map{
			"name": name,
			"age":  age,
		})
		return nil
	}
}
