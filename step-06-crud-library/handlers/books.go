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
			respondError(w, http.StatusInternalServerError, err.Error())
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(books)
	}
}
func GetBookByID(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		path := strings.TrimPrefix(r.URL.Path, "/books/")
		id, err := strconv.Atoi(path)
		if err != nil {
			respondError(w, http.StatusBadRequest, "invalid book id")
			return
		}

		err = db.Get(&book, "SELECT * FROM books WHERE id = $1", id)
		if err != nil {
			if err == sql.ErrNoRows {
				respondError(w, http.StatusNotFound, "book not found")
				return
			}
			respondError(w, http.StatusInternalServerError, err.Error())
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
			respondError(w, http.StatusBadRequest, "invalid JSON")
			return
		}

		if err := newBook.Validate(); err != nil {
			respondError(w, http.StatusBadRequest, err.Error())
			return
		}

		query := `
			INSERT INTO books (title, author, year)
			VALUES ($1, $2, $3)
			RETURNING id, title, author, year, created_at;
		`
		err = db.Get(&newBook, query, newBook.Title, newBook.Author, newBook.Year)
		if err != nil {
			respondError(w, http.StatusInternalServerError, err.Error())
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newBook)
	}
}
func UpdateBook(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		path := strings.TrimPrefix(r.URL.Path, "/books/")
		id, err := strconv.Atoi(path)
		if err != nil {
			respondError(w, http.StatusBadRequest, "invalid book id")
			return
		}

		if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
			respondError(w, http.StatusBadRequest, "invalid JSON")
			return
		}

		if err := book.Validate(); err != nil {
			respondError(w, http.StatusBadRequest, err.Error())
			return
		}

		query := `
			UPDATE books
			SET title = $1, author = $2, year = $3
			WHERE id = $4
			RETURNING id, title, author, year, created_at;
		`

		err = db.Get(&book, query, book.Title, book.Author, book.Year, id)
		if err != nil {
			if err == sql.ErrNoRows {
				respondError(w, http.StatusNotFound, "book not found")
				return
			}
			respondError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(book)
	}
}
func DeleteBook(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimPrefix(r.URL.Path, "/books/")
		id, err := strconv.Atoi(path)
		if err != nil {
			respondError(w, http.StatusBadRequest, "invalid book id")
			return
		}

		result, err := db.Exec("DELETE FROM books WHERE id = $1", id)
		if err != nil {
			respondError(w, http.StatusInternalServerError, err.Error())
			return
		}

		rows, _ := result.RowsAffected()
		if rows == 0 {
			respondError(w, http.StatusNotFound, "book not found")
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
