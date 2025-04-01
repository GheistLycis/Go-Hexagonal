package main

import (
	cli "Go-Hexagonal/cmd/cli"
	tcp "Go-Hexagonal/cmd/tcp"
	web "Go-Hexagonal/cmd/web"
	postgres "Go-Hexagonal/infra/postgres"
	"log"
	"os"
)

var isBuild = "false"
var entry = "web"

func main() {
	args := os.Args

	if isBuild == "false" {
		entry = args[1]
		args = args[2:]
	} else {
		args = args[1:]
	}

	switch entry {
	case "web":
		log.Print("RUNNING WEB SERVER")
		web.Init(postgres.Init())
	case "cli":
		log.Print("RUNNING CLI PROGRAM")
		cli.Init(args)
	case "tcp":
		log.Print("RUNNING TCP SERVER")
		tcp.Init()
	default:
		log.Fatal("ARGUMENT FOR ENTRYPOINT NOT SUPPORTED -", entry)
	}
}
