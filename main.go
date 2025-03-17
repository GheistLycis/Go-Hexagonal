package main

import (
	tcp "Go-Hexagonal/cmd/tcp"
	web "Go-Hexagonal/cmd/web"
	db "Go-Hexagonal/infra/db"
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
