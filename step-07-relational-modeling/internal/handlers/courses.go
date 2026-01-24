package handlers

import (
	"encoding/json"
	"net/http"
	"step-07-relational-modeling/internal/models"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

func CreateCourse(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newCourse models.Course

		if err := json.NewDecoder(r.Body).Decode(&newCourse); err != nil {
			http.Error(w, "invalid JSON body", http.StatusBadRequest)
			return
		}
		if newCourse.Title == "" || newCourse.Description == "" {
			http.Error(w, "Title and Description are required", http.StatusBadRequest)
			return
		}
		query := `
			INSERT INTO courses (title, description)
			VALUES ($1, $2)
			RETURNING id, title, description, created_at;
		`
		err := db.Get(
			&newCourse,
			query,
			newCourse.Title,
			newCourse.Description,
		)

		if err != nil {
			http.Error(w, "failed to create course", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newCourse)
	}
}

func ListCourses(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		courses := []models.Course{}

		err := db.Select(&courses, "SELECT * FROM courses ORDER BY id")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(courses)
	}
}

func ListStudentsByCourse(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		courseID := chi.URLParam(r, "course_id")
		if courseID == "" {
			http.Error(w, "course_id is required", http.StatusBadRequest)
			return
		}

		students := []models.Student{}

		query := `
			SELECT s.*
			FROM students s
			INNER JOIN enrollments e ON e.student_id = s.id
			WHERE e.course_id = $1
			ORDER BY s.id;
		`

		err := db.Select(&students, query, courseID)
		if err != nil {
			http.Error(w, "failed to list students", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(students)
	}
}
