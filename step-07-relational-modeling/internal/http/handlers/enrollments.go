package handlers

import (
	"encoding/json"
	"net/http"
	"step-07-relational-modeling/internal/models"

	"github.com/jmoiron/sqlx"
)

func CreateEnrollment(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var enrollment models.Enrollment

		if err := json.NewDecoder(r.Body).Decode(&enrollment); err != nil {
			http.Error(w, "invalid JSON body", http.StatusBadRequest)
			return
		}

		if enrollment.StudentID == 0 || enrollment.CourseID == 0 {
			http.Error(w, "student_id and course_id are required", http.StatusBadRequest)
			return
		}

		query := `
			INSERT INTO enrollments (student_id, course_id)
			VALUES ($1, $2)
			RETURNING student_id, course_id, enrolled_at;
		`

		err := db.Get(
			&enrollment,
			query,
			enrollment.StudentID,
			enrollment.CourseID,
		)

		if err != nil {
			http.Error(w, "failed to create enrollment", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(enrollment)
	}
}
