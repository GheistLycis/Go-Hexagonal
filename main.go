package main

import (
	"Go-Hexagonal/cmd/api"
	db "Go-Hexagonal/infra/db"
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
