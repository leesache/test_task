package models

import (
	"database/sql"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	var err error
	var dbSQL *sql.DB

	retries := 10
	for i := 0; i < retries; i++ {
		dbSQL, err = sql.Open("pgx", dsn)
		if err == nil {
			err = dbSQL.Ping()
			if err == nil {
				break
			}
		}
		log.Printf("Database not ready, retrying in 5 seconds... (%d/%d)", i+1, retries)
		time.Sleep(3 * time.Second)
	}

	if err != nil {
		log.Fatalf("Failed to connect to database after %d retries: %v", retries, err)
	}

	DB, err = gorm.Open(postgres.New(postgres.Config{Conn: dbSQL}), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to initialize GORM: %v", err)
	}

	log.Println("Database connection established successfully!")
}
