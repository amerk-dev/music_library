package main

import (
	"log"
	"test/internal/frameworks"
)

func main() {
	log.Println("Hello World")
	server := frameworks.NewServer()
	server.Start()
}
