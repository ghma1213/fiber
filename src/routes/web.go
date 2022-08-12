package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	// "github.com/gofibe"
)

func Router() *fiber.App {
	// App 생성, 템플릿 설정
	app := fiber.New(fiber.Config{
		Views: html.New("./views", "html"),
	})

	// Route 설정
	// app.Get("/", controllers.Index)
	return app
}
