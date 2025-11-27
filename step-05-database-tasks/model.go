package main

import "time"

type Task struct {
	ID        int       `db:"id" json:"id"`
	Title     string    `db:"title" json:"title"`
	Completed bool      `db:"completed" json:"completed"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}
