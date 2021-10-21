package config

import (
	"log"

	"github.com/DungBuiTien1999/udemy-api/internal/validator"
)

type AppConfig struct {
	InfoLog      *log.Logger
	ErrorLog     *log.Logger
	InProduction bool
	Validator    *validator.Validation
}
