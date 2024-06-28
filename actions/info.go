package actions

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gobuffalo/buffalo"
)

type userResponse struct {
	Name       string `json:"name" default:"Иван"`
	Surname    string `json:"surname" default:"Иванов"`
	Patronymic string `json:"patronymic" default:"Иванович"`
	Address    string `json:"address" default:"г. Москва, ул. Ленина, д. 5, кв. 1"`
}

func MyHandler(w http.ResponseWriter, r *http.Request) {
	user := userResponse{
		Name:       "Иван",
		Surname:    "Иванов",
		Patronymic: "Иванович",
		Address:    "г. Москва, ул. Ленина, д. 5, кв. 1",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func GetUserByPassportDumb(c buffalo.Context) error {
	var err error

	passSerParam := c.Param("passport_serie")
	passSer, err := strconv.Atoi(passSerParam)
	if err != nil {
		return c.Render(http.StatusBadRequest, r.JSON("Invalid passport serie"))
	}
	// must be 4-digits number
	if passSer < 1000 || passSer > 9999 {
		return c.Render(http.StatusBadRequest, r.JSON("Invalid passport serie"))
	}

	passNumbParam := c.Param("passport_number")
	passNumb, err := strconv.Atoi(passNumbParam)
	if err != nil {
		return c.Render(http.StatusBadRequest, r.JSON("Invalid passport number"))
	}
	// must be 6-digits number
	if passNumb < 1000 || passNumb > 999999 {
		return c.Render(http.StatusBadRequest, r.JSON("Invalid passport number"))
	}

	user := userResponse{
		Name:       "Иван",
		Surname:    "Иванов",
		Patronymic: "Иванович",
		Address:    "г. Москва, ул. Ленина, д. 5, кв. 1",
	}
	return c.Render(http.StatusOK, r.JSON(user))
}
