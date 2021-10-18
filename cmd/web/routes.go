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

	mux.Get("/users", handlers.Repo.AllUsers)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
