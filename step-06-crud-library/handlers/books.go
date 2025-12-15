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

func RegisterBook(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newBook models.Book

		err := json.NewDecoder(r.Body).Decode(&newBook)
		if err != nil {
			http.Error(w, "invalid JSON", http.StatusBadRequest)
			return
		}

		query := `
			INSERT INTO books (title, author, year)
			VALUES ($1, $2, $3)
			RETURNING id, title, author, year, created_at;
		`
		err = db.Get(&newBook, query, newBook.Title, newBook.Author, newBook.Year)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newBook)
	}
}
