package db

import (
	"database/sql"
	"fmt"
	"log"
	"message-service/pkg/config"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB(cfg *config.Config) {
	var err error
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("Error connecting to database: %q", err)
	}

	log.Println("Successfully connected to the database")

	// Создание таблицы, если она не существует
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS messages (
		id SERIAL PRIMARY KEY,
		content TEXT NOT NULL,
		is_processed BOOLEAN DEFAULT false,
		processed_at TIMESTAMP
	);`
	_, err = DB.Exec(createTableQuery)
	if err != nil {
		log.Fatalf("Error initializing database table: %q", err)
	}

	log.Println("Database and table are initialized successfully.")
}
