package main

import (
	"log"
	"net/http"
	"step-06-crud-library/db"
	"step-06-crud-library/handlers"
)

func main() {
	database := db.ConnectDB()
	defer database.Close()

	http.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		handlers.ListBooks(database)(w, r)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
	log.Println("Server running on http://localhost:8080")
}
