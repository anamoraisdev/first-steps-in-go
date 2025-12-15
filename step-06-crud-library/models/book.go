package models

import "time"

type Book struct {
	ID        int       `db:"id" json:"id"`
	Title     string    `db:"title" json:"title"`
	Author    string    `db:"author" json:"author"`
	Year      int       `db:"year" json:"year"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}
