package database

import (
	"database/sql"
	"log"
	"os"
)

func InitDB() *sql.DB {

	connStr := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME")
	DB, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	return DB
}
