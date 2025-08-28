package routes

import (
	"fiber-app/controllers"
	"github.com/gofiber/fiber/v2"
)



func SetupRoutes(app *fiber.App) {

	app.Get("/about", func(c *fiber.Ctx) error {
		return c.Render("about", fiber.Map{
			"Title": "About • Fiber App",
		}, "layout")
	})

	app.Get("/users/:id", controllers.UserView)


}





