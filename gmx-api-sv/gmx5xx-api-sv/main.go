package main

import (
	"log"
)

func main() {
	store, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}
	server := Gmx5xxAPIServer("0.0.0.0:3000", store)
	server.Run()
}
