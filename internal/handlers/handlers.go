package handlers

import (
	"net/http"

	"github.com/DungBuiTien1999/udemy-api/internal/config"
	"github.com/DungBuiTien1999/udemy-api/internal/driver"
	"github.com/DungBuiTien1999/udemy-api/internal/helpers"
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
		helpers.ServerError(w, err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		resp := models.GenericError{
			Message: err.Error(),
		}
		helpers.ToJSON(resp, w)
		return
	}

	resp := models.JSONResponse{
		Status: http.StatusOK,
		Data:   users,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	helpers.ToJSON(resp, w)
}

func (m *Repository) TestPostRequest(w http.ResponseWriter, r *http.Request) {

	var t models.User
	err := helpers.FromJSON(&t, r.Body)
	if err != nil {
		panic(err)
	}
	errs := m.App.Validator.Validate(t)
	if len(errs) != 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		resp := models.ValidationError{
			Messages: errs.Errors(),
		}
		helpers.ToJSON(resp, w)
		return
	}
}
