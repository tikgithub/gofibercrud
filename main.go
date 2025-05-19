package main

import (
	"demo/db"
	"demo/routes"
	"log"

	fiber "github.com/gofiber/fiber/v2"
)

func main() {
	log.Println("Starting server...")
	err := db.InitDb()
	db.Migrate()

	if err != nil {
		log.Fatal("Error initializing database:", err)
	}
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "ok",
			"message": "Welcome to the server!",
		})
	})

	routes.SetupRoutes(app)

	err = app.Listen(":3000")

	if err != nil {
		log.Println("Error starting server:", err)
	} else {
		log.Println("Server started on port 3000")
	}
}
