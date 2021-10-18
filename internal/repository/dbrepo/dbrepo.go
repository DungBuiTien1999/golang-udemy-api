package dbrepo

import (
	"database/sql"

	"github.com/DungBuiTien1999/udemy-api/internal/config"
	"github.com/DungBuiTien1999/udemy-api/internal/repository"
)

type mysqlDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewMySQLRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &mysqlDBRepo{
		App: a,
		DB:  conn,
	}
}
