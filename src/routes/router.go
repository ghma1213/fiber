package routes

import (
	"fiber/src/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func Router() *fiber.App {
	// Create a new engine
	engine := html.New("./views", ".html")

	// Pass the engine to the Views
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Route 설정
	app.Get("/", controllers.Index)
	return app
}
