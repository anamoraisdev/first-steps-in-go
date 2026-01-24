package handlers

import (
	"encoding/json"
	"net/http"
	"step-07-relational-modeling/internal/models"

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
