package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/anamoraisdev/first-steps-in-go/step-05-database-tasks/models"
	"github.com/jmoiron/sqlx"
)

func ListTasks(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tasks := []models.Task{}

		err := db.Select(&tasks, "SELECT * FROM tasks ORDER BY id")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tasks)
	}
}

func CreateTask(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input struct {
			Title string `json:"title"`
		}

		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			http.Error(w, "invalid JSON", http.StatusBadRequest)
			return
		}
		if input.Title == "" {
			http.Error(w, "title is required", http.StatusBadRequest)
			return
		}

		var newTask models.Task
		query := `
			INSERT INTO tasks (title)
			VALUES ($1)
			RETURNING id, title, completed, created_at;
		`
		err = db.Get(&newTask, query, input.Title)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newTask)
	}
}
