package routes

import (
	"github.com/bartalcorn/apistarter/person"
	"github.com/gofiber/fiber/v2"
)

// Post displays the default page via POST
func Post() fiber.Handler {
	return func(c *fiber.Ctx) error {
		person := person.Person{}
		_ = c.BodyParser(person)
		// if err != nil {
		// 	c.Status(503).Send([]byte(err.Error()))
		// 	return err
		// }
		_ = c.JSON(person)

		return nil
	}
}
