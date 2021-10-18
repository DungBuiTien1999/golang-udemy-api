package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/DungBuiTien1999/udemy-api/internal/config"
	"github.com/DungBuiTien1999/udemy-api/internal/driver"
	"github.com/DungBuiTien1999/udemy-api/internal/models"
	"github.com/DungBuiTien1999/udemy-api/internal/repository"
	"github.com/DungBuiTien1999/udemy-api/internal/repository/dbrepo"
)

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
	DB  repository.DatabaseRepo
}

// Repo the repository used by the handlers
var Repo *Repository

// NewRepo create a new repository
func NewRepo(a *config.AppConfig, db *driver.DB) *Repository {
	return &Repository{
		App: a,
		DB:  dbrepo.NewMySQLRepo(db.SQL, a),
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) AllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := m.DB.AllUsers()
	if err != nil {
		resp := models.JSONUser{
			Status: http.StatusBadRequest,
			Data:   users,
		}
		out, _ := json.MarshalIndent(resp, "", "     ")
		w.Header().Set("Content-Type", "application/json")
		w.Write(out)
		return
	}

	resp := models.JSONUser{
		Status: http.StatusOK,
		Data:   users,
	}

	out, _ := json.MarshalIndent(resp, "", "     ")

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}
