package main

import (
	"log"
)

func main() {
	store, err := NewSqliteStorage("todo.db")

	if err != nil {
		log.Fatal("Error initializing database: ", err)
	}

	server := NewAPIServer("localhost:8080", store)
	err = server.Run()

	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
