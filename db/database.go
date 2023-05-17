package db

import (
	"database/sql"
	"log"
	"os"
	
	- "github.com/database-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal('Error laoding .env file')
	}
}

func InitDB() {

	loadEnv();

	connStr := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME");
	db, err := sql.Open("mysql", connStr);
	if err != nil {
		log.Fatal("Error connecting to the database:", err)	
	}
	DB = db
}