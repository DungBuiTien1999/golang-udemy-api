package handlers

import (
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/DungBuiTien1999/udemy-api/internal/config"
	"github.com/DungBuiTien1999/udemy-api/internal/driver"
	"github.com/DungBuiTien1999/udemy-api/internal/helpers"
	"github.com/DungBuiTien1999/udemy-api/internal/models"
	"github.com/DungBuiTien1999/udemy-api/internal/repository"
	"github.com/DungBuiTien1999/udemy-api/internal/repository/dbrepo"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/joho/godotenv/autoload"
	"github.com/thanhpk/randstr"
	"golang.org/x/crypto/bcrypt"
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

// TasksOfUser returns json which contains slice of task of an user
func (m *Repository) TasksOfUser(w http.ResponseWriter, r *http.Request) {
	exploted := strings.Split(r.RequestURI, "/")
	userID, err := strconv.Atoi(exploted[3])
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

	tasks, err := m.DB.GetTasksByUserID(userID)
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
		Data:   tasks,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	helpers.ToJSON(resp, w)
}

// Register create an account for an user
func (m *Repository) Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := helpers.FromJSON(&user, r.Body)
	if err != nil {
		panic(err)
	}
	errs := m.App.Validator.Validate(user)
	if len(errs) != 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		resp := models.ValidationError{
			Messages: errs.Errors(),
		}
		helpers.ToJSON(resp, w)
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	newUser := models.User{
		Username: user.Username,
		Password: string(hashedPassword),
	}

	err = m.DB.InsertUser(newUser)
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
		Status:  http.StatusCreated,
		Message: "Registed user successfully!",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	helpers.ToJSON(resp, w)
}

// AddTask creates a task of an user into database
func (m *Repository) AddTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task

	err := helpers.FromJSON(&task, r.Body)
	if err != nil {
		panic(err)
	}
	errs := m.App.Validator.Validate(task)
	if len(errs) != 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		resp := models.ValidationError{
			Messages: errs.Errors(),
		}
		helpers.ToJSON(resp, w)
		return
	}

	err = m.DB.InsertTask(task)
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
		Status:  http.StatusCreated,
		Message: "Created task successfully!",
		Data:    nil,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	helpers.ToJSON(resp, w)
}

// UpdateTask updates a task of an user
func (m *Repository) UpdateTask(w http.ResponseWriter, r *http.Request) {
	exploted := strings.Split(r.RequestURI, "/")
	id, err := strconv.Atoi(exploted[3])
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

	var task models.Task

	err = helpers.FromJSON(&task, r.Body)
	if err != nil {
		panic(err)
	}
	errs := m.App.Validator.Validate(task)
	if len(errs) != 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		resp := models.ValidationError{
			Messages: errs.Errors(),
		}
		helpers.ToJSON(resp, w)
		return
	}

	newTask, err := m.DB.UpdateTaskByID(id, task)
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
		Status:  http.StatusOK,
		Message: "Updated task successfully!",
		Data:    newTask,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	helpers.ToJSON(resp, w)
}

// Authentication check if user authorized and return access token and refresh token
func (m *Repository) Authentication(w http.ResponseWriter, r *http.Request) {
	var u models.User
	err := helpers.FromJSON(&u, r.Body)
	if err != nil {
		panic(err)
	}
	errs := m.App.Validator.Validate(u)
	if len(errs) != 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		resp := models.ValidationError{
			Messages: errs.Errors(),
		}
		helpers.ToJSON(resp, w)
		return
	}

	user, err := m.DB.GetUserByUsername(u.Username)
	if err != nil || user.Username == "" {
		helpers.ServerError(w, err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		resp := models.AuthenticationResp{
			Authorized:   false,
			Messages:     "Username doesn't exist",
			AccessToken:  "",
			RefreshToken: "",
		}
		helpers.ToJSON(resp, w)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		resp := models.AuthenticationResp{
			Authorized:   false,
			Messages:     "Password wrong",
			AccessToken:  "",
			RefreshToken: "",
		}
		helpers.ToJSON(resp, w)
		return
	} else if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		resp := models.AuthenticationResp{
			Authorized:   false,
			Messages:     err.Error(),
			AccessToken:  "",
			RefreshToken: "",
		}
		helpers.ToJSON(resp, w)
		return
	}

	rfToken := randstr.String(32)
	err = m.DB.UpdateRefreshToken(user.ID, rfToken)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		resp := models.AuthenticationResp{
			Authorized:   false,
			Messages:     err.Error(),
			AccessToken:  "",
			RefreshToken: "",
		}
		helpers.ToJSON(resp, w)
		return
	}

	secretToken := os.Getenv("SECRET_TOKEN")

	payload := models.NewPayload(user.ID, time.Hour*24) // access token has expired at 24 hours
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	accessToken, err := at.SignedString([]byte(secretToken))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		resp := models.AuthenticationResp{
			Authorized:   false,
			Messages:     err.Error(),
			AccessToken:  "",
			RefreshToken: "",
		}
		helpers.ToJSON(resp, w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resp := models.AuthenticationResp{
		Authorized:   true,
		Messages:     "Login successfully",
		AccessToken:  accessToken,
		RefreshToken: rfToken,
	}
	helpers.ToJSON(resp, w)
}
