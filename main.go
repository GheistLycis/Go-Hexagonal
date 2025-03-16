package main

import (
	db "Go-Hexagonal/adapters/db"
	"Go-Hexagonal/cmd/api"
	"Go-Hexagonal/cmd/tcp"
)

func main() {
	go api.Init(8000, db.Init(true))
	go tcp.Init(8080)

	select {}
}
