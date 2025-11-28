package main

import (
	"log"
	"net/http"

	"github.com/anamoraisdev/first-steps-in-go/step-05-database-tasks/db"
	"github.com/anamoraisdev/first-steps-in-go/step-05-database-tasks/handlers"
)

func main() {
	db := db.ConnectDB()
	defer db.Close()

	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.ListTasks(db)(w, r)
		case http.MethodPost:
			handlers.CreateTask(db)(w, r)
		}
	})

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
