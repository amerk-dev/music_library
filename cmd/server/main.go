package main

import (
	"test/internal/frameworks"
	"test/internal/store"
)

func main() {
	store.DataBaseInit()

	server := frameworks.NewServer()
	server.Start()
}
