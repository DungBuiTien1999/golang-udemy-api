package handlers

import (
	"log"
	"os"
	"testing"

	"github.com/DungBuiTien1999/udemy-api/internal/config"
	"github.com/DungBuiTien1999/udemy-api/internal/validator"
)

var app config.AppConfig

func TestMain(m *testing.M) {
	app.InProduction = false

	app.Validator = validator.NewValidation()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	repo := NewTestingRepo(&app)
	NewHandlers(repo)

	os.Exit(m.Run())
}

// func getRoutes() http.Handler {
// 	mux := chi.NewMux()

// 	mux.Use(middleware.Recoverer)

// 	mux.Route("/api/auth", func(r chi.Router) {
// 		r.Post("/", Repo.Authentication)
// 		r.Post("/refresh", Repo.PostRefreshToken)
// 	})
// 	mux.Route("/api/users", func(r chi.Router) {
// 		r.Post("/", Repo.Register)
// 	})
// 	mux.Route("/api/tasks", func(r chi.Router) {
// 		r.Get("/{userID}", Repo.TasksOfUser)
// 		r.Post("/", Repo.AddTask)
// 		r.Put("/{id}", Repo.UpdateTask)
// 	})

// 	fileServer := http.FileServer(http.Dir("./static/"))
// 	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

// 	return mux
// }
