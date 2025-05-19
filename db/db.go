package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDb() error {
	var err error
	dsn := "host=69.10.53.189 user=appuser password=123456 dbname=myapp port=5432 sslmode=disable"
	DB, err = sql.Open("postgres", dsn)

	if err != nil {
		log.Println("Error connecting to database:", err)
		return err
	}

	if err = DB.Ping(); err != nil {
		log.Println("Error pinging database:", err)
		return err
	}
	log.Println("Connected to database")
	return nil

}
