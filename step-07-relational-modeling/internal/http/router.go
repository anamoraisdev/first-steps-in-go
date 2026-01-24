package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"

	"step-07-relational-modeling/internal/http/handlers"
)

func NewRouter(db *sqlx.DB) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	r.Route("/students", func(r chi.Router) {
		r.Post("/", handlers.RegisterStudent(db))
		r.Get("/", handlers.ListStudents(db))
	})

	r.Route("/courses", func(r chi.Router) {
		r.Post("/", handlers.CreateCourse(db))
		r.Get("/", handlers.ListCourses(db))

		r.Route("/{course_id}", func(r chi.Router) {
			r.Post("/lessons", handlers.CreateLesson(db))
			r.Get("/lessons", handlers.ListLessonsByCourse(db))
		})
	})

	return r
}
