package main

import (
	"log"
	"net/http"
	"step-08-transactions/internal/db"
)

func main() {
	database := db.ConnectDB()
	defer database.Close()

	log.Println("🚀 Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
