package routes

import (
	"fiber-app/database"
	"fiber-app/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func SetupRoutes(app *fiber.App) {

	// Home page
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "🚀 Fiber App is running!",
			"Name":  "Anand",
		}, "layout")
	})
}



