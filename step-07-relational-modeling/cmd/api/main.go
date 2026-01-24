package main

import (
	"log"
	"net/http"
	"step-07-relational-modeling/internal/db"
	httpRouter "step-07-relational-modeling/internal/http"
)

func main() {
	database := db.ConnectDB()
	defer database.Close()

	router := httpRouter.NewRouter(database)

	log.Println("🚀 Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
