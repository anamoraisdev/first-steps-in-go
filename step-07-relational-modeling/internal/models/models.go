package models

import "time"

type Student struct {
	ID        int       `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	Email     string    `db:"email" json:"email"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

type Course struct {
	ID          int       `db:"id" json:"id"`
	Title       string    `db:"title" json:"title"`
	Description string    `db:"description" json:"description"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
}

type Lesson struct {
	ID        int       `db:"id" json:"id"`
	CourseID  int       `db:"course_id" json:"course_id"`
	Title     string    `db:"title" json:"title"`
	StartsAt  time.Time `db:"starts_at" json:"starts_at"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

type Enrollment struct {
	StudentID  int       `db:"student_id" json:"student_id"`
	CourseID   int       `db:"course_id" json:"course_id"`
	EnrolledAt time.Time `db:"enrolled_at" json:"enrolled_at"`
}
