package seed

import (
	"log"
	"math/rand"

	"github.com/jmoiron/sqlx"
)

func SeedBooks(db *sqlx.DB, amount int) {

	titles := []string{
		"O Hobbit", "Clean Code", "1984", "Dom Quixote",
		"Arquitetura Limpa", "Go em Ação",
	}

	authors := []string{
		"Tolkien", "Robert Martin", "George Orwell",
	}

	for i := 0; i < amount; i++ {
		title := titles[rand.Intn(len(titles))]
		author := authors[rand.Intn(len(authors))]
		year := rand.Intn(2024-1950) + 1950

		_, err := db.Exec(`
			INSERT INTO books (title, author, year)
			VALUES ($1, $2, $3)
		`, title, author, year)

		if err != nil {
			log.Fatal(err)
		}
	}
}
