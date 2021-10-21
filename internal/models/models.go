package models

import "time"

// User is the user model
type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username" validate:"required"`
	Password  string    `json:"password" validate:"required,passwd"`
	RfToken   string    `json:"rf_token"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Task is the task model
type Task struct {
	ID        int       `json:"id"`
	Title     string    `json:"title" validate:"required"`
	Complete  int       `json:"complete"`
	UserID    int       `json:"user_id"  validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
