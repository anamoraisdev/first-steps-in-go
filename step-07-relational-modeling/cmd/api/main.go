package main

import (
	"log"
	"net/http"
	"step-07-relational-modeling/internal/db"
	"step-07-relational-modeling/internal/handlers"
)

func main() {
	database := db.ConnectDB()
	defer database.Close()

	http.HandleFunc("/students", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handlers.RegisterStudent(database)(w, r)
		case http.MethodGet:
			handlers.ListStudents(database)(w, r)
		}
	})

	http.HandleFunc("/courses", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handlers.CreateCourse(database)(w, r)
		case http.MethodGet:
			handlers.ListCourses(database)(w, r)
		}
	})

	log.Println("🚀 Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
