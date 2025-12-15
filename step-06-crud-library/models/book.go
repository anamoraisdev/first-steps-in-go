package models

import "time"

type Book struct {
	ID        int       `db:"id" json:"id"`
	Title     string    `db:"title" json:"title"`
	Author    string    `json:"author"`
	Year      int       `json:"year"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}
