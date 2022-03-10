package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"toru/backend/db"
	"toru/backend/handlers"
)


func Router() {
	db.ConnectDB()
	
    app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	app.Get("/",func (c *fiber.Ctx) error {
		return c.SendString("Connected Server API")
	})
	app.Get("/products",handlers.ShowProduct)
	app.Post("/products/add",handlers.SaveProduct)

    app.Listen(":3001")
}