package main

import (
	db "Go-Hexagonal/adapters/db"
	"Go-Hexagonal/cmd/tcp"
	"Go-Hexagonal/cmd/web"
)

func main() {
	go web.Init(8000, db.Init(true))
	go tcp.Init(8080)

	select {}
}
