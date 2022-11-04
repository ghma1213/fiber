package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/v1/:value", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World! " + c.Params("value"))
	})

	app.Get("/v1/:name?", func(c *fiber.Ctx) error {
		if c.Params("name") != "" {
			return c.SendString("Hello " + c.Params("name"))
			// => Hello john
		}
		return c.SendString("Where is ㅇㅇㅇ?")
	})

	app.Get("/api/*", func(c *fiber.Ctx) error {
		return c.SendString("API path: " + c.Params("*"))
		// => API path: user/john
	})
	app.Static("/", "./public")
	app.Listen(":3000")
}
