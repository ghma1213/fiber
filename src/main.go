package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	handler := func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	}

	handler2 := func(c *fiber.Ctx) error {
		return c.SendString("hihi")
	}

	api := app.Group("/api") // /api

	v1 := api.Group("/v1", func(c *fiber.Ctx) error { // middleware for /api/v1
		c.Set("Version", "v1")
		return c.Next()
	})

	v2 := api.Group("/v2", func(c *fiber.Ctx) error { // middleware for /api/v1
		c.Set("Version", "v2")
		return c.Next()
	})

	v1.Get("/list", handler) // /api/v1/list
	v1.Get("/user", handler) // /api/v1/user

	v2.Get("/hi", handler2)

	app.Listen(":3000")
}
