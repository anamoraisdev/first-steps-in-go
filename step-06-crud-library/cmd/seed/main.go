package main

import (
	"fmt"

	"step-06-crud-library/internal/db"
	"step-06-crud-library/internal/seed"
)

func main() {
	database := db.ConnectDB()
	defer database.Close()

	seed.SeedBooks(database, 30)

	fmt.Println("📚 Database seeded successfully")
}
