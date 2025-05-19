package db

import (
	"fmt"
	"log"
)

func Migrate() error {
	query := `
	CREATE TABLE IF NOT EXISTS hydrogreen.users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		email VARCHAR(100) NOT NULL UNIQUE);
	`

	_, err := DB.Exec(query)

	if err != nil {
		log.Println("Error creating table:", err)
		return err
	}

	fmt.Println("Table created successfully")
	return nil
}
