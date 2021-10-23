package helpers

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"runtime/debug"
	"time"

	"github.com/DungBuiTien1999/udemy-api/internal/config"
	"github.com/DungBuiTien1999/udemy-api/internal/models"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/joho/godotenv/autoload"
)

var app *config.AppConfig

var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

// NewHelpers sets up app config for helpers
func NewHelpers(a *config.AppConfig) {
	app = a
}

func ClientError(w http.ResponseWriter, status int) {
	app.InfoLog.Println("Client error with status of", status)
	http.Error(w, http.StatusText(status), status)
}

func ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLog.Println(trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func VerifyToken(tokenString string) (*models.Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(os.Getenv("SECRET_TOKEN")), nil
	}
	jwtToken, err := jwt.ParseWithClaims(tokenString, &models.Payload{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}
		return nil, ErrExpiredToken
	}

	payload, ok := jwtToken.Claims.(*models.Payload)
	if !ok {
		return nil, ErrInvalidToken
	}

	return payload, nil
}

func CreateNewAccessToken(userID int) (string, error) {
	secretToken := os.Getenv("SECRET_TOKEN")

	payload := models.NewPayload(userID, time.Hour*24) // access token has expired at 24 hours
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	accessToken, err := at.SignedString([]byte(secretToken))
	if err != nil {
		return "", err
	}
	return accessToken, nil
}
