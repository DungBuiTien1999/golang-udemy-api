package models

type JSONUser struct {
	Status int    `json:"status"`
	Data   []User `json:"data"`
}
