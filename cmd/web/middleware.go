package main

import (
	"errors"
	"net/http"
	"os"

	"github.com/DungBuiTien1999/udemy-api/internal/helpers"
	"github.com/DungBuiTien1999/udemy-api/internal/models"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/joho/godotenv/autoload"
)

var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

// Auth protects roots which needs authenticated
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := verifyToken(r)
		if err != nil {
			helpers.ServerError(w, err)
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

		next.ServeHTTP(w, r)
	})
}

func verifyToken(r *http.Request) (*models.Payload, error) {
	tokenString := r.Header.Get("x-access-token")
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
