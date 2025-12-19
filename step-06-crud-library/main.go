package main

import (
	"log"
	"net/http"
	"step-06-crud-library/internal/db"
	"step-06-crud-library/internal/handlers"
	"step-06-crud-library/internal/middlewares"
)

func main() {
	database := db.ConnectDB()
	defer database.Close()

	mux := http.NewServeMux()

	mux.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.ListBooks(database)(w, r)
		case http.MethodPost:
			handlers.RegisterBook(database)(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/books/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetBookByID(database)(w, r)
		case http.MethodPut:
			handlers.UpdateBook(database)(w, r)
		case http.MethodDelete:
			handlers.DeleteBook(database)(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	handler := middlewares.Logger(
		middlewares.JSON(mux),
	)

	log.Fatal(http.ListenAndServe(":8080", handler))
}
