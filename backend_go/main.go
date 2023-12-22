package main

import (
	"log"
    "os"
)

func main() {

    d := os.Getenv("DB_PATH")

    if d == "" {
        d = "todo.db"
    }

    store, err := NewSqliteStorage(d)

	if err != nil {
		log.Fatal("Error initializing database: ", err)
	}

    i := os.Getenv("APP_INTERFACE")

    if i == "" {
        i = ":8080"
    }

    server := NewAPIServer(i, store)
	err = server.Run()

	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
