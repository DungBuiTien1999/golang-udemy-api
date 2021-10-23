package main

import (
	"net/http"

	"github.com/DungBuiTien1999/udemy-api/internal/config"
	"github.com/DungBuiTien1999/udemy-api/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewMux()

	mux.Use(middleware.Recoverer)
	mux.Use(cors.AllowAll().Handler)

	mux.Route("/api/auth", func(r chi.Router) {
		r.Post("/", handlers.Repo.Authentication)
		r.Post("/refresh", handlers.Repo.PostRefreshToken)
	})
	mux.Route("/api/users", func(r chi.Router) {
		r.Post("/", handlers.Repo.Register)
	})
	mux.Route("/api/tasks", func(r chi.Router) {
		r.Use(Auth)
		r.Get("/{userID}", handlers.Repo.TasksOfUser)
		r.Post("/", handlers.Repo.AddTask)
		r.Put("/{id}", handlers.Repo.UpdateTask)
	})

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
