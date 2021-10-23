package main

import (
	"net/http"

	"github.com/DungBuiTien1999/udemy-api/internal/helpers"
	"github.com/DungBuiTien1999/udemy-api/internal/models"
	_ "github.com/joho/godotenv/autoload"
)

// Auth protects roots which needs authenticated
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("x-access-token")
		_, err := helpers.VerifyToken(tokenString)
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
