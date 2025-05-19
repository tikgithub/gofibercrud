package main

import (
	"log"

	fiber "github.com/gofiber/fiber/v2"
)

func main() {
	log.Println("Starting server...")
	app := fiber.New()
	
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "ok",
			"message": "Welcome to the server!",
		})
	})

	err := app.Listen(":3000")

	if err != nil {
		log.Println("Error starting server:", err)
	} else {
		log.Println("Server started on port 3000")

	}
}
