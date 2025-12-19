package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"step-06-crud-library/internal/models"
	"step-06-crud-library/internal/utils"
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"
)

type PaginatedResponse struct {
	Data  []models.Book `json:"data"`
	Page  int           `json:"page"`
	Limit int           `json:"limit"`
	Total int           `json:"total"`
}

func ListBooks(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		books := []models.Book{}

		page := 1
		limit := 20

		pageStr := r.URL.Query().Get("page")
		limitStr := r.URL.Query().Get("limit")

		if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
			page = p
		}

		if l, err := strconv.Atoi(limitStr); err == nil && l > 0 && l <= 100 {
			limit = l
		}

		offset := (page - 1) * limit
		var total int

		err := db.Get(&total, "SELECT COUNT(*) FROM books")
		if err != nil {
			utils.RespondError(w, http.StatusInternalServerError, "failed to count books")
			return
		}

		query := `
			SELECT * FROM books
			ORDER BY id
			LIMIT $1 OFFSET $2
		`

		err = db.Select(&books, query, limit, offset)
		if err != nil {
			utils.RespondError(w, http.StatusInternalServerError, "failed to fetch books")
			return
		}
		response := PaginatedResponse{
			Data:  books,
			Page:  page,
			Limit: limit,
			Total: total,
		}

		json.NewEncoder(w).Encode(response)
	}
}
func GetBookByID(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		path := strings.TrimPrefix(r.URL.Path, "/books/")
		id, err := strconv.Atoi(path)
		if err != nil {
			utils.RespondError(w, http.StatusBadRequest, "invalid book id")
			return
		}

		err = db.Get(&book, "SELECT * FROM books WHERE id = $1", id)
		if err != nil {
			if err == sql.ErrNoRows {
				utils.RespondError(w, http.StatusNotFound, "book not found")
				return
			}
			utils.RespondError(w, http.StatusInternalServerError, err.Error())
			return
		}

		json.NewEncoder(w).Encode(book)
	}
}
func RegisterBook(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newBook models.Book

		err := json.NewDecoder(r.Body).Decode(&newBook)
		if err != nil {
			utils.RespondError(w, http.StatusBadRequest, "invalid JSON")
			return
		}

		if err := newBook.Validate(); err != nil {
			utils.RespondError(w, http.StatusBadRequest, err.Error())
			return
		}

		query := `
			INSERT INTO books (title, author, year)
			VALUES ($1, $2, $3)
			RETURNING id, title, author, year, created_at;
		`
		err = db.Get(&newBook, query, newBook.Title, newBook.Author, newBook.Year)
		if err != nil {
			utils.RespondError(w, http.StatusInternalServerError, err.Error())
			return
		}

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
			utils.RespondError(w, http.StatusBadRequest, "invalid book id")
			return
		}

		if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
			utils.RespondError(w, http.StatusBadRequest, "invalid JSON")
			return
		}

		if err := book.Validate(); err != nil {
			utils.RespondError(w, http.StatusBadRequest, err.Error())
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
				utils.RespondError(w, http.StatusNotFound, "book not found")
				return
			}
			utils.RespondError(w, http.StatusInternalServerError, err.Error())
			return
		}

		json.NewEncoder(w).Encode(book)
	}
}
func DeleteBook(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimPrefix(r.URL.Path, "/books/")
		id, err := strconv.Atoi(path)
		if err != nil {
			utils.RespondError(w, http.StatusBadRequest, "invalid book id")
			return
		}

		result, err := db.Exec("DELETE FROM books WHERE id = $1", id)
		if err != nil {
			utils.RespondError(w, http.StatusInternalServerError, err.Error())
			return
		}

		rows, _ := result.RowsAffected()
		if rows == 0 {
			utils.RespondError(w, http.StatusNotFound, "book not found")
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
