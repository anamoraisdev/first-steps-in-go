package main

import (
	"log"
	"net/http"
)

func main() {
	db := ConnectDB()
	defer db.Close()

	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			ListTasks(db)(w, r)
		case http.MethodPost:
			CreateTask(db)(w, r)
		}
	})

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
