package routesv1

import (
	"github.com/bartalcorn/apistarter/person"
	"github.com/gofiber/fiber/v2"
)

// Person is a sample route for POSTing data
func Person() fiber.Handler {
	return func(c *fiber.Ctx) error {
		person := person.Person{}
		err := c.BodyParser(person)
		if err != nil {
			return err
		}
		return nil
	}
}
