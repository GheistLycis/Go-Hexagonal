package main

import (
	db "Go-Hexagonal/adapters/db"
	"Go-Hexagonal/cmd/tcp"
	"Go-Hexagonal/cmd/web"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file - %v", err)
	}

	go web.Init(db.Init())
	go tcp.Init()

	select {}
}
