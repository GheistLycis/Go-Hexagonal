package main

import (
	"Go-Hexagonal/cmd/api"
	"Go-Hexagonal/cmd/event"
	"fmt"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("RECOVERED", r)
		}
	}()

	api.Init(8000)
	event.Init()
}
