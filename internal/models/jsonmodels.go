package models

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JSONResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

// ValidationError is a collection of validation error messages
type ValidationError struct {
	Messages []string `json:"messages"`
}

type AuthenticationResp struct {
	Authorized   bool   `json:"authorized"`
	Messages     string `json:"messages"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type AuthCustomClaims struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}

type Payload struct {
	UserID    int       `json:"user_id"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func NewPayload(userID int, duration time.Duration) *Payload {
	payload := &Payload{
		UserID:    userID,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload
}

func (p *Payload) Valid() error {
	if time.Now().After(p.ExpiredAt) {
		return errors.New("token has expired")
	}
	return nil
}
