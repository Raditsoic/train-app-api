package main

import (
	"os"

	"github.com/Raditsoic/train-app-api/controller"
	"github.com/Raditsoic/train-app-api/database"
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	err := database.Connect()
	if err != nil {
		panic(err)
	}

	err = database.Migrate()
	if err != nil {
		panic(err)
	}
	
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to train-app-api")
	})

	controller.SetupRoutes(app)

	app.Listen(os.Getenv("PORT"))
}
