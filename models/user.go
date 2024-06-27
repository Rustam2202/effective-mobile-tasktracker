package models

import "time"

type User struct {
	ID             int       `json:"id" db:"id"`
	Surname        string    `json:"surname" db:"surname"`
	Name           string    `json:"name" db:"name"`
	Patronymic     string    `json:"patronymic" db:"patronymic"`
	Address        string    `json:"address" db:"address"`
	PassportSerie  int       `json:"passport_serie" db:"passport_serie"`
	PassportNumber int       `json:"passport_number" db:"passport_number"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
}
