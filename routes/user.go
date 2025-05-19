package routes

import (
	"demo/db"
	"demo/models"
	"log"

	"github.com/gofiber/fiber/v2"
)

func GetAllUsers(c *fiber.Ctx) error {

	/* Check Database Connection is OK */
	if db.DB == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Database connection is not established",
		})
	}

	/* Query Builder */
	rows, err := db.DB.Query("SELECT * FROM hydrogreen.users")
	if err != nil {
		log.Println("Error querying database:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Error querying database " + err.Error(),
		})
	}

	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			log.Println("Error scanning row:", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":  "error",
				"message": "Error scanning row " + err.Error(),
			})
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error iterating rows:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Error iterating rows " + err.Error(),
		})
	}

	return c.JSON(users)

}

func SetupRoutes(app *fiber.App) {
	// Define your routes here
	app.Get("/users", GetAllUsers)
}
