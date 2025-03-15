package db

import (
	"fmt"
	"log"
	"os"

	user "Go-Hexagonal/infra/db/user"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

/*
Init generates the DSN string from local .env and connects to the DB.

-doMigration: whether to migrate every mapped model.
*/
func Init(doMigration bool) {
	dsn := getDSN()

	DB = connect(dsn)

	if doMigration {
		migrate(DB)
	}
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
	err := db.AutoMigrate(&user.UserModel{})

	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
}
