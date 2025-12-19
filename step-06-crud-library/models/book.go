package models

import (
	"errors"
	"strings"
	"time"
)

type Book struct {
	ID        int       `db:"id" json:"id"`
	Title     string    `db:"title" json:"title"`
	Author    string    `db:"author" json:"author"`
	Year      int       `db:"year" json:"year"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

func (b *Book) Validate() error {
	if strings.TrimSpace(b.Title) == "" {
		return errors.New("title is required")
	}

	if strings.TrimSpace(b.Author) == "" {
		return errors.New("author is required")
	}

	currentYear := time.Now().Year()
	if b.Year < 1000 || b.Year > currentYear {
		return errors.New("invalid year")
	}

	return nil
}
