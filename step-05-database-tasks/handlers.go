package main

import (
	"encoding/json"
	"net/http"

	"github.com/jmoiron/sqlx"
)

func ListTasks(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tasks := []Task{}

		err := db.Select(&tasks, "SELECT * FROM tasks ORDER BY id")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tasks)
	}
}
