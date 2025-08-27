package main

import (
	"log"

	"fiber-app/models"
	"fiber-app/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	database.Connect()

	// Table create (AutoMigrate)
	database.DB.AutoMigrate(&models.User{})

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Home route
	app.Get("/", func(c *fiber.Ctx) error {
		var users []models.User
	    database.DB.Find(&users)


		return c.Render("index", fiber.Map{
			"Title":     "Home • Fiber App",
			"PageTitle": "Dashboard",
			"Name":      "Anand",
			"Year":      2025,
			"Users":     users,
		}, "layout")
	})


	// Form page
	app.Get("/create", func(c *fiber.Ctx) error {
		return c.Render("create", fiber.Map{
			"Title": "Create New User",
		}, "layout")
	})

    

	app.Post("/create", func(c *fiber.Ctx) error {
			user := new(models.User)

			user.Name = c.FormValue("name")
			user.Email = c.FormValue("email")
			user.Password = c.FormValue("password")

			result := database.DB.Create(&user)
			if result.Error != nil {
				return c.Status(500).SendString(result.Error.Error())
			}

			return c.Redirect("/create?success=1")
	})


	log.Fatal(app.Listen(":3000"))
}
