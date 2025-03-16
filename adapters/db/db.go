package db

import (
	"fmt"
	"log"
	"os"

	user "Go-Hexagonal/adapters/db/user"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

/*
Init generates the DSN string from local .env, connects to the DB and returns the active connection.

-m: whether to auto-migrate every mapped model.
*/
func Init(m bool) *gorm.DB {
	dsn := getDSN()
	DB := connect(dsn)

	if m {
		migrate(DB)
	}

	return DB
}

func getDSN() string {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)

		return ""
	}

	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
}

func connect(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)

		return nil
	}

	return db
}

func migrate(db *gorm.DB) {
	if err := db.AutoMigrate(&user.UserModel{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
}
