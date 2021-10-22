package repository

import "github.com/DungBuiTien1999/udemy-api/internal/models"

type DatabaseRepo interface {
	AllUsers() ([]models.User, error)
	GetTasksByUserID(userID int) ([]models.Task, error)
	InsertTask(task models.Task) error
	UpdateTaskByID(id int, task models.Task) (models.Task, error)
	GetUserByUsername(username string) (models.User, error)
	InsertUser(user models.User) error
	UpdateRefreshToken(id int, rfToken string) error
	IsValidRefreshToken(id int, rfToken string) bool
}
