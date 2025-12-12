package main

import (
	"log"

	"library-api/db"
)

func main() {
	database := db.ConnectDB()
	defer database.Close()

	log.Println("Server running on http://localhost:8080")
}
