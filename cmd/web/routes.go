package main

import (
	"net/http"

	"github.com/DungBuiTien1999/udemy-api/internal/config"
	"github.com/DungBuiTien1999/udemy-api/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewMux()

	mux.Use(middleware.Recoverer)

	mux.Route("/api/auth", func(r chi.Router) {
		r.Post("/", handlers.Repo.TestPostRequest)
		r.Post("/refresh", handlers.Repo.TestPostRequest)
	})
	mux.Route("/api/users", func(r chi.Router) {
		r.Post("/", handlers.Repo.TestPostRequest)
	})
	mux.Route("/api/tasks", func(r chi.Router) {
		r.Get("/{userID}", handlers.Repo.AllUsers)
		r.Post("/", handlers.Repo.TestPostRequest)
		r.Patch("/{id}", handlers.Repo.TestPostRequest)

	})

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
