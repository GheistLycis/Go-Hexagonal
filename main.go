package main

import (
	db "Go-Hexagonal/adapters/db"
	"Go-Hexagonal/cmd/api"
	"Go-Hexagonal/cmd/tcp"
	"fmt"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("RECOVERED", r)
		}
	}()

	dbConn := db.Init(true)
	api.Init(8000, dbConn)
	tcp.Init(8080)
}
