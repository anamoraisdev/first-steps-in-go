package handlers

import (
	"encoding/json"
	"net/http"
	"step-06-crud-library/models"

	"github.com/jmoiron/sqlx"
)

func ListBooks(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		books := []models.Book{}

		err := db.Select(&books, "SELECT * FROM books ORDER BY id")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(books)
	}
}
