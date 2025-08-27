package controllers

import (
    "fiber-app/database"
    "fiber-app/models"
    "github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
    user := new(models.User)

    user.Name = c.FormValue("name")
    user.Email = c.FormValue("email")
    user.Password = c.FormValue("password")

	
    result := database.DB.Create(&user)
    if result.Error != nil {
        return c.Status(500).SendString(result.Error.Error())
    }

    return c.SendString("✅ User Saved Successfully")
}
