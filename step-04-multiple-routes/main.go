package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Balance float64 `json:"balance"`
}

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type Order struct {
	ID       int     `json:"id"`
	UserID   int     `json:"user_id"`
	Products []int   `json:"products"`
	Total    float64 `json:"total"`
}

var users = []User{}
var nextUserID = 1

func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(users)
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	var u User

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "Invalid user JSON", http.StatusBadRequest)
		return
	}

	u.ID = nextUserID
	nextUserID++

	users = append(users, u)

	json.NewEncoder(w).Encode(u)
}

func main() {
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getUsersHandler(w, r)

		case http.MethodPost:
			createUserHandler(w, r)

		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Server running on http://localhost:8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server error:", err)
	}
}
