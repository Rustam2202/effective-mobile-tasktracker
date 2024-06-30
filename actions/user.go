package actions

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"tasktracker/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/pop/v6"
)

// @Summary Get all users
// @Description Get a list of all users with optional filters for pagination and search
// @Tags Users
// @Accept json
// @Produce json
// @Param id query integer false "User ID"
// @Param name query string false "User name"
// @Param surname query string false "User surname"
// @Param patronymic query string false "User patronymic"
// @Param passportSerie query integer false "User passport serie"
// @Param passportNumber query integer false "User passport number"
// @Param address query string false "User address"
// @Param page query integer true "Page number" default(0)
// @Param perPage query integer true "Number of users per page" default(10)
// @Success 200 {object} models.User
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Router /users [get]
func GetAllUsers(c buffalo.Context) error {
	users := []models.User{}
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return c.Render(http.StatusInternalServerError, r.JSON("Internal server error"))
	}
	query := tx.Q()

	if c.Param("id") != "" {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.Render(http.StatusBadRequest, r.JSON("Invalid id"))
		}
		if id < 0 {
			return c.Render(http.StatusBadRequest, r.JSON("Invalid id"))
		}
		query = query.Where("id = ?", id)
	}
	if c.Param("name") != "" {
		query = query.Where("name = ?", c.Param("name"))
	}
	if c.Param("surname") != "" {
		query = query.Where("surname = ?", c.Param("surname"))
	}
	if c.Param("patronymic") != "" {
		query = query.Where("patronymic = ?", c.Param("patronymic"))
	}
	if c.Param("passportSerie") != "" {
		passSer, err := strconv.Atoi(c.Param("passportSerie"))
		if err != nil {
			return c.Render(http.StatusBadRequest, r.JSON("Invalid passport serie"))
		}
		if passSer < 0 || passSer > 4 {
			return c.Render(http.StatusBadRequest, r.JSON("Invalid passport serie"))
		}
		query = query.Where("passport_serie = ?", passSer)
	}
	if c.Param("passportNumber") != "" {
		passNumb, err := strconv.Atoi(c.Param("passportNumber"))
		if err != nil {
			return c.Render(http.StatusBadRequest, r.JSON("Invalid passport number"))
		}
		if passNumb < 0 || passNumb > 6 {
			return c.Render(http.StatusBadRequest, r.JSON("Invalid passport number"))
		}
		query = query.Where("passport_number = ?", passNumb)
	}
	if c.Param("address") != "" {
		query = query.Where("address = ?", c.Param("address"))
	}

	if c.Param("page") == "" {
		return c.Render(http.StatusBadRequest, r.JSON("Invalid perPage number"))
	}
	page, err := strconv.Atoi(c.Param("page"))
	if err != nil {
		return c.Render(http.StatusBadRequest, r.JSON("Invalid page number"))
	}
	if page < 0 {
		return c.Render(http.StatusBadRequest, r.JSON("Invalid page number"))
	}
	if c.Param("per_page") == "" {
		return c.Render(http.StatusBadRequest, r.JSON("Invalid perPage number"))
	}
	perPage, err := strconv.Atoi(c.Param("per_page"))
	if err != nil {
		return c.Render(http.StatusBadRequest, r.JSON("Invalid perPage number"))
	}
	if perPage < 0 {
		return c.Render(http.StatusBadRequest, r.JSON("Invalid perPage number"))
	}

	err = query.All(&users)
	if err != nil {
		return c.Render(http.StatusNotFound, r.JSON("Users not found"))
	}
	return c.Render(http.StatusOK, r.JSON(users))
}

type createUserRequest struct {
	PassportNuber string `json:"passportNumber"`
}

type responseUser struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Address    string `json:"address"`
}

// @Summary Create a new user
// @Description Create a new user with the provided information
// @Tags Users
// @Accept json
// @Produce json
// @Param userReq body createUserRequest true "User request body"
// @Success 200 {string} string
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /users [post]
func CreateUser(c buffalo.Context) error {
	var (
		err            error
		userReq        createUserRequest
		passportSerie  int
		passportNumber int
	)
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return c.Render(http.StatusInternalServerError, r.JSON("Internal server error"))
	}
	err = c.Bind(&userReq)
	if err != nil {
		return c.Render(http.StatusBadRequest, r.JSON("Invalid request"))
	}

	splitNumber := strings.Split(userReq.PassportNuber, " ")
	if len(splitNumber) != 2 {
		return c.Render(http.StatusBadRequest, r.JSON("Invalid passport number"))
	}
	passportSerie, err = strconv.Atoi(splitNumber[0])
	if err != nil {
		return c.Render(http.StatusBadRequest, r.JSON("Invalid passport serie"))
	}
	passportNumber, err = strconv.Atoi(splitNumber[1])
	if err != nil {
		return c.Render(http.StatusBadRequest, r.JSON("Invalid passport number"))
	}

	envy.Load("../.env")
	infoURL, err := envy.MustGet("INFO_URL")
	if err != nil {
		return c.Render(http.StatusInternalServerError, r.JSON("Internal server error"))
	}

	req := fmt.Sprintf(
		"%s/info?passportSerie=%s&passportNumber=%s",
		infoURL, fmt.Sprint(passportSerie), fmt.Sprint(passportNumber))

	resp, err := http.Get(req)
	if err != nil {
		return c.Render(http.StatusInternalServerError, r.JSON("Internal server error"))
	}

	user := responseUser{}
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return c.Render(http.StatusInternalServerError, r.JSON("Internal server error"))
	}
	err = json.Unmarshal(bodyBytes, &user)
	if err != nil {
		return c.Render(http.StatusInternalServerError, r.JSON("Internal server error"))
	}
	err = tx.Create(&models.User{
		Surname:        user.Surname,
		Name:           user.Name,
		Patronymic:     user.Patronymic,
		Address:        user.Address,
		PassportSerie:  passportSerie,
		PassportNumber: passportNumber,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	})

	if err != nil {
		return c.Render(http.StatusInternalServerError, r.JSON("Internal server error"))
	}

	return c.Render(http.StatusOK, r.JSON("User created"))
}

type updateUserRequest struct {
	ID             int       `json:"id"`
	Surname        string    `json:"surname"`
	Name           string    `json:"name"`
	Patronymic     string    `json:"patronymic"`
	Address        string    `json:"address"`
	PassportSerie  int       `json:"passport_serie"`
	PassportNumber int       `json:"passport_number"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// @Summary Update a user
// @Description Update an existing user with the provided information
// @Tags Users
// @Accept json
// @Produce json
// @Param user_id path integer true "User ID"
// @Param userReq body updateUserRequest true "User request body"
// @Success 200 {object} models.User
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Router /users/{user_id} [put]
func UpdateUser(c buffalo.Context) error {
	var userReq updateUserRequest
	err := c.Bind(&userReq)
	if err != nil {
		return c.Render(http.StatusBadRequest, r.JSON("Invalid request"))
	}

	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return c.Render(http.StatusInternalServerError, r.JSON("Internal server error"))
	}

	user := &models.User{}
	if err := tx.Find(user, c.Param("user_id")); err != nil {
		return c.Render(http.StatusBadRequest, r.JSON("Invalid request"))
	}

	user.Name = userReq.Name
	user.Surname = userReq.Surname
	user.Patronymic = userReq.Patronymic
	user.PassportSerie = userReq.PassportSerie
	user.PassportNumber = userReq.PassportNumber
	user.Address = userReq.Address
	user.CreatedAt = userReq.CreatedAt
	user.UpdatedAt = time.Now()

	if err := models.DB.Update(&user); err != nil {
		return c.Render(http.StatusInternalServerError, r.JSON("Internal server error"))
	}

	return c.Render(http.StatusOK, r.JSON(user))
}

// @Summary Delete a user
// @Description Delete an existing user by ID
// @Tags Users
// @Accept json
// @Produce json
// @Param user_id path integer true "User ID"
// @Success 200 {string} string
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Router /users/{user_id} [delete]
func DeleteUser(c buffalo.Context) error {
	id, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		return c.Render(http.StatusBadRequest, r.JSON("Invalid id"))
	}

	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return c.Render(http.StatusInternalServerError, r.JSON("Internal server error"))
	}

	err = tx.Destroy(&models.User{ID: id})
	if err != nil {
		return c.Render(http.StatusNotFound, r.JSON("User not found"))
	}

	return c.Render(http.StatusOK, r.JSON("User deleted"))
}
