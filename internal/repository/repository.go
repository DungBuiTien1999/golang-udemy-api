package repository

import "github.com/DungBuiTien1999/udemy-api/internal/models"

type DatabaseRepo interface {
	AllUsers() ([]models.User, error)
}
