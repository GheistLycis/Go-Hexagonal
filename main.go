package main

import (
	cli "Go-Hexagonal/cmd/cli"
	tcp "Go-Hexagonal/cmd/tcp"
	web "Go-Hexagonal/cmd/web"
	db "Go-Hexagonal/infra/db"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var entry = "web"

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file - %v", err)
	}
}

func main() {
	if args := os.Args; len(args) >= 2 {
		entry = args[1]
	}

	switch entry {
	case "web":
		log.Print("RUNNING WEB SERVER")
		web.Init(db.Init())
	case "cli":
		log.Print("RUNNING CLI PROGRAM")
		cli.Init()
	case "tcp":
		log.Print("RUNNING TCP SERVER")
		tcp.Init()
	default:
		log.Fatal("ARGUMENT FOR ENTRYPOINT NOT SUPPORTED -", entry)
	}
}
