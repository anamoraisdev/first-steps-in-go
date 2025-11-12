package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Message struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
}

var messages []Message
var nextID = 1

func getMessagesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
}

func createMessageHandler(w http.ResponseWriter, r *http.Request) {
	var newMessage Message
	err := json.NewDecoder(r.Body).Decode(&newMessage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	newMessage.ID = nextID
	nextID++
	messages = append(messages, newMessage)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newMessage)

}
func main() {
	http.HandleFunc("/messages", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			getMessagesHandler(w, r)
		} else if r.Method == http.MethodPost {
			createMessageHandler(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
