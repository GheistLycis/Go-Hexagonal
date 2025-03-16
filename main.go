package main

import (
	db "Go-Hexagonal/adapters/db"
	"Go-Hexagonal/cmd/api"
	"fmt"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("RECOVERED", r)
		}
	}()

	db.Init(true)
	api.Init(8000)
}
