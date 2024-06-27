package models

type User struct {
	ID             int    `json:"id" db:"id"`
	Surname        string `json:"surname" db:"surname"`
	Name           string `json:"name" db:"name"`
	Patronymic     string `json:"patronymic" db:"patronymic"`
	Address        string `json:"address" db:"address"`
	PassportSerie  int    `json:"passport_serie" db:"passport_serie"`
	PassportNumber int    `json:"passport_number" db:"passport_number"`
}
