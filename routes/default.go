package routes

import (
	// "fmt"

	"github.com/gofiber/fiber/v2"
)

var html = `
<html>
<head>
  <title>Https Test</title>
</head>
<body style="padding:24px">
  <h2>There's no need to fear,</h2>
  <h1 style="color:#336699;">Bartman is here!</h1>
</body>
</html>
`

func Default() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Format([]byte(html))
	}
}
