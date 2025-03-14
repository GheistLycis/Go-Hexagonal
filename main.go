package main

import (
	"Go-Hexagonal/cmd/api"
	"Go-Hexagonal/cmd/event"
)

func main() {
	api.Init(8000)
	event.Init()
}
