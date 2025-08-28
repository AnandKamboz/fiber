package controllers

import (
    "fmt"            
    "fiber-app/database"
    "fiber-app/models"
    
    "github.com/gofiber/fiber/v2"
)


func UserView(c *fiber.Ctx) error {
	id := c.Params("id")
	fmt.Println("User ID:", id)

	var user models.User
	result := database.DB.First(&user, id)
	if result.Error != nil {
		return c.Status(404).SendString("User not found")
	}

	return c.Render("view", fiber.Map{
		"Title": "User View • Fiber App",
		"User":  user,
	}, "layout")
}

// func DeleteUser(c *fiber.Ctx) error {
// 	id :c.Params("id")
// 	var user models.User
// 	result := database.DB.First(&user, id)

// }


func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := database.DB.Delete(&models.User{}, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(204) 
}
