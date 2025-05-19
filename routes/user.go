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

func AddUser(c *fiber.Ctx) error{

	user := new(models.User);
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Error parsing request body " + err.Error(),
		})
	}

	query := `INSERT INTO hydrogreen.users (name, email) VALUES ($1, $2) RETURNING id`;
	err := db.DB.QueryRow(query, user.Name, user.Email).Scan(&user.ID);
	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Error inserting user " + err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "User created successfully",
		"user":    user,});
}

func RemoveUser(c *fiber.Ctx) error{
	id := c.Params("id");
	query :=`DELETE FROM hydrogreen.users WHERE id = $1`;
	_, err := db.DB.Exec(query, id);
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Error deleting user " + err.Error(),
		})
	}
	
	log.Println("ID: ", id);
	return c.SendString("User deleted successfully");
}

func SetupRoutes(app *fiber.App) {
	// Define your routes here
	app.Get("/users", GetAllUsers)
	app.Post("/users", AddUser)
	app.Delete("/users/:id", RemoveUser)
}
