package handlers

import (
	"encoding/json"
	"net/http"
	"step-07-relational-modeling/internal/models"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

func CreateLesson(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newLesson models.Lesson

		courseIDParam := chi.URLParam(r, "course_id")
		courseID, err := strconv.Atoi(courseIDParam)
		if err != nil {
			http.Error(w, "invalid course_id", http.StatusBadRequest)
			return
		}

		if err := json.NewDecoder(r.Body).Decode(&newLesson); err != nil {
			http.Error(w, "invalid JSON body", http.StatusBadRequest)
			return
		}

		if newLesson.Title == "" {
			http.Error(w, "title is required", http.StatusBadRequest)
			return
		}

		query := `
			INSERT INTO lessons (course_id, title, starts_at)
			VALUES ($1, $2, $3)
			RETURNING id, course_id, title, starts_at, created_at;
		`

		err = db.Get(
			&newLesson,
			query,
			courseID,
			newLesson.Title,
			newLesson.StartsAt,
		)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newLesson)
	}
}
