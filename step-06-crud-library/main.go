package main

import (
	"log"
	"net/http"
	"step-06-crud-library/db"
	"step-06-crud-library/handlers"
)

func main() {
	database := db.ConnectDB()
	defer database.Close()

	http.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.ListBooks(database)(w, r)
		case http.MethodPost:
			handlers.RegisterBook(database)(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)

		}
	})

	http.HandleFunc("/books/", func(w http.ResponseWriter, r *http.Request) {
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
	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
