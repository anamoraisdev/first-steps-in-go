package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"step-06-crud-library/models"
	"strconv"
	"strings"

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
func GetBookByID(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimPrefix(r.URL.Path, "/books/")
		id, err := strconv.Atoi(path)
		if err != nil {
			http.Error(w, "invalid book id", http.StatusBadRequest)
			return
		}

		var book models.Book
		err = db.Get(&book, "SELECT * FROM books WHERE id = $1", id)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "book not found", http.StatusNotFound)
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(book)
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
