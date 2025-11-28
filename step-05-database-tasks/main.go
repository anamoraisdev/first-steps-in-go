package main

import (
	"log"
	"net/http"
)

func main() {
	db := ConnectDB()
	defer db.Close()

	http.HandleFunc("/tasks", ListTasks(db))
	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
