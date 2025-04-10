package postgres

import (
	"fmt"
	"log"
	"os"
	"strconv"

	user "Go-Hexagonal/src/user/infra"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

/*
Init generates the DSN string from .env and connects to the database.

Returns the active connection.

-m: whether to auto-migrate every mapped model.
*/
func Init() *gorm.DB {
	dsn := getDSN()
	conn := connect(dsn)

	autoMigrateDB, err := strconv.ParseBool(os.Getenv("DB_AUTO_MIGRATE"))
	if err != nil {
		log.Fatal("[DB] Error while parsing ENV variable DB_AUTO_MIGRATE -", err)
	}
	if autoMigrateDB {
		migrate(conn)
	}

	return conn
}

func getDSN() string {
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
		log.Fatalf("[DB] Failed to connect - %v", err)

		return nil
	}

	return db
}

func migrate(db *gorm.DB) {
	if err := db.AutoMigrate(&user.UserModel{}); err != nil {
		log.Fatalf("[DB] Failed to migrate - %v", err)
	}

	log.Print("[DB] Automigration done successfully")
}
