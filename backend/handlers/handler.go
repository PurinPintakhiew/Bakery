package handlers

import (
	"toru/backend/db"
	"toru/backend/model"
	"github.com/gofiber/fiber/v2"
)

func ShowProduct(c *fiber.Ctx) error {
	var data []model.Product
	db.DB.Find(&data)
	return c.JSON(data)
}

func SaveProduct(c *fiber.Ctx) error {
    data := new(model.Product)

    err := c.BodyParser(data)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
    }
    err = db.DB.Create(&data).Error
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create note", "data": err})
    }
	return c.JSON(fiber.Map{"message": "Save Completed"})
}