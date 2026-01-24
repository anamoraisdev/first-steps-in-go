package handlers

import (
	"encoding/json"
	"net/http"
	"step-07-relational-modeling/internal/models"

	"github.com/jmoiron/sqlx"
)

func RegisterStudent(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newStudent models.Student

		if err := json.NewDecoder(r.Body).Decode(&newStudent); err != nil {
			http.Error(w, "invalid JSON body", http.StatusBadRequest)
			return
		}
		if newStudent.Name == "" || newStudent.Email == "" {
			http.Error(w, "name and email are required", http.StatusBadRequest)
			return
		}
		query := `
			INSERT INTO students (name, email)
			VALUES ($1, $2)
			RETURNING id, name, email, created_at;
		`
		err := db.Get(
			&newStudent,
			query,
			newStudent.Name,
			newStudent.Email,
		)

		if err != nil {
			http.Error(w, "failed to create student", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newStudent)
	}
}

func ListStudents(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		students := []models.Student{}

		err := db.Select(&students, "SELECT * FROM students ORDER BY id")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(students)
	}
}
