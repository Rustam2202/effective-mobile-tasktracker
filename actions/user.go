package actions

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	log "tasktracker/logger"
	"tasktracker/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/pop/v6"
)

// GetAllUsers
//
//	@Summary		Get all users
//	@Description	Get a list of all users with optional filters for pagination and search
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			id				query		integer	false	"User ID"
//	@Param			name			query		string	false	"User name"
//	@Param			surname			query		string	false	"User surname"
//	@Param			patronymic		query		string	false	"User patronymic"
//	@Param			passportSerie	query		integer	false	"User passport serie"
//	@Param			passportNumber	query		integer	false	"User passport number"
//	@Param			address			query		string	false	"User address"
//	@Param			page			query		integer	true	"Page number"				default(1)
//	@Param			perPage			query		integer	true	"Number of users per page"	default(10)
//	@Success		200				{array}		models.User
//	@Failure		400				{string}	string	"Invalid request"
//	@Failure		404				{string}	string	"Users not found"
//	@Router			/user [get]
func GetAllUsers(c buffalo.Context) error {
	var (
		err     error
		users   = []models.User{}
		page    int
		perPage int
	)

	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		log.Logger.Error().Msg("Start transaction error")
		return c.Render(http.StatusInternalServerError, r.JSON("Internal server error"))
	}
	query := tx.Q()

	if c.Param("id") != "" {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Logger.Error().Msg("Invalid id")
			return c.Render(http.StatusBadRequest, r.JSON("Invalid id"))
		}
		query = query.Where("id = ?", id)
		log.Logger.Debug().Msg("ID added to query")
	}
	if c.Param("name") != "" {
		query = query.Where("name = ?", c.Param("name"))
		log.Logger.Debug().Msg("Name added to query")
	}
	if c.Param("surname") != "" {
		query = query.Where("surname = ?", c.Param("surname"))
		log.Logger.Debug().Msg("Surname added to query")
	}
	if c.Param("patronymic") != "" {
		query = query.Where("patronymic = ?", c.Param("patronymic"))
		log.Logger.Debug().Msg("Patronymic added to query")
	}
	if c.Param("passportSerie") != "" {
		passSer, err := strconv.Atoi(c.Param("passportSerie"))
		if err != nil {
			return c.Render(http.StatusBadRequest, r.JSON("Invalid passport serie"))
		}
		if passSer < 1000 || passSer > 9999 {
			log.Logger.Error().Msg("passportSerie must contain 4 digits")
			return c.Render(http.StatusBadRequest, r.JSON("Invalid passport serie"))
		}
		query = query.Where("passport_serie = ?", passSer)
		log.Logger.Debug().Msg("Passport serie added to query")
	}
	if c.Param("passportNumber") != "" {
		passNumb, err := strconv.Atoi(c.Param("passportNumber"))
		if err != nil {
			log.Logger.Error().Msg("failed parsing passportNumber")
			return c.Render(http.StatusBadRequest, r.JSON("Invalid passport number"))
		}
		if passNumb < 1000 || passNumb > 999999 {
			log.Logger.Error().Msg("passportNumber must contain 6 digits")
			return c.Render(http.StatusBadRequest, r.JSON("Invalid passport number"))
		}
		query = query.Where("passport_number = ?", passNumb)
		log.Logger.Debug().Msg("Passport number added to query")
	}
	if c.Param("address") != "" {
		query = query.Where("address = ?", c.Param("address"))
		log.Logger.Debug().Msg("Address added to query")
	}

	if c.Param("page") == "" {
		page = 1
		log.Logger.Debug().Msg("Page number is empty, set to 1")
	} else {
		page, err = strconv.Atoi(c.Param("page"))
		if err != nil {
			log.Logger.Error().Msg("failed to parse page number")
			return c.Render(http.StatusBadRequest, r.JSON("Invalid page number"))
		}
		if page < 0 {
			page = 1
			log.Logger.Debug().Msg("Page number is less than 0, set to 1")
		}
	}
	if c.Param("per_page") == "" {
		perPage = 10
		log.Logger.Debug().Msg("perPage is empty, set to 10")
	} else {
		perPage, err = strconv.Atoi(c.Param("per_page"))
		if err != nil {
			log.Logger.Error().Msg("failed to parse perPage number")
			return c.Render(http.StatusBadRequest, r.JSON("Invalid perPage number"))
		}
		if perPage < 0 {
			perPage = 10
			log.Logger.Debug().Msg("perPage is less than 0, set to 10")
		}
	}

	err = query.Paginate(page, perPage).All(&users)
	if err != nil {
		log.Logger.Error().Err(err).Msg("Failed to get users")
		return c.Render(http.StatusNotFound, r.JSON("Users not found"))
	}
	log.Logger.Debug().Msg("Successfully get users")
	return c.Render(http.StatusOK, r.JSON(users))
}

type createUserRequest struct {
	PassportNuber string `json:"passportNumber"`
}

type responseUser struct {
	People struct {
		Name       string `json:"name"`
		Surname    string `json:"surname"`
		Patronymic string `json:"patronymic"`
		Address    string `json:"address"`
	} `json:"people"`
}

// CreateUser
//
//	@Summary		Create a new user
//	@Description	Create a new user with the provided information
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			userReq	body		createUserRequest	true	"User request body"
//	@Success		200		{string}	string				"User created"
//	@Failure		400		{string}	string				"Invalid request"
//	@Failure		500		{string}	string				"Internal server error"
//	@Router			/user [post]
func CreateUser(c buffalo.Context) error {
	var (
		err            error
		userReq        createUserRequest
		passportSerie  int
		passportNumber int
	)
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		log.Logger.Error().Msg("Start transaction error")
		return c.Render(http.StatusInternalServerError, r.JSON("Internal server error"))
	}
	err = c.Bind(&userReq)
	if err != nil {
		log.Logger.Error().Err(err).Msg("Failed to bind request")
		return c.Render(http.StatusBadRequest, r.JSON("Invalid request"))
	}

	splitNumber := strings.Split(userReq.PassportNuber, " ")
	if len(splitNumber) != 2 {
		log.Logger.Error().Msg("passportNumber must contain 2 parts separated by space")
		return c.Render(http.StatusBadRequest, r.JSON("Invalid passport number"))
	}
	passportSerie, err = strconv.Atoi(splitNumber[0])
	if err != nil {
		log.Logger.Error().Msg("failed parsing passportSerie")
		return c.Render(http.StatusBadRequest, r.JSON("Invalid passport serie"))
	}
	passportNumber, err = strconv.Atoi(splitNumber[1])
	if err != nil {
		log.Logger.Error().Msg("failed parsing passportNumber")
		return c.Render(http.StatusBadRequest, r.JSON("Invalid passport number"))
	}

	envy.Load("../.env")
	infoURL, err := envy.MustGet("INFO_URL")
	if err != nil {
		log.Logger.Error().Msg("Failed to get INFO_URL from .env")
		return c.Render(http.StatusInternalServerError, r.JSON("Internal server error"))
	}

	req := fmt.Sprintf(
		"%s/info?passportSerie=%s&passportNumber=%s",
		infoURL, fmt.Sprint(passportSerie), fmt.Sprint(passportNumber))

	resp, err := http.Get(req)
	if err != nil {
		log.Logger.Error().Err(err).Msg("Failed to get user from external service")
		return c.Render(http.StatusInternalServerError, r.JSON("Internal server error"))
	}

	user := responseUser{}
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Logger.Error().Err(err).Msg("Failed to read response body")
		return c.Render(http.StatusInternalServerError, r.JSON("Internal server error"))
	}
	err = json.Unmarshal(bodyBytes, &user)
	if err != nil {
		log.Logger.Error().Err(err).Msg("Failed to unmarshal response body")
		return c.Render(http.StatusInternalServerError, r.JSON("Internal server error"))
	}
	err = tx.Create(&models.User{
		Surname:        user.People.Surname,
		Name:           user.People.Name,
		Patronymic:     user.People.Patronymic,
		Address:        user.People.Address,
		PassportSerie:  passportSerie,
		PassportNumber: passportNumber,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	})

	if err != nil {
		log.Logger.Error().Err(err).Msg("Failed to create user")
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

// UpdateUser
//
//	@Summary		Update a user
//	@Description	Update an existing user with the provided information
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			user_id	path		integer				true	"User ID"
//	@Param			userReq	body		updateUserRequest	true	"User request body"
//	@Success		200		{object}	models.User
//	@Failure		400		{string}	string	"Invalid request"
//	@Failure		404		{string}	string	"User not found"
//	@Router			/user [put]
func UpdateUser(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		log.Logger.Error().Msg("Start transaction error")
		return c.Render(http.StatusInternalServerError, r.JSON("Internal server error"))
	}

	var userReq updateUserRequest
	if err := c.Bind(&userReq); err != nil {
		log.Logger.Error().Err(err).Msg("Failed to bind request with user struct")
		return c.Render(http.StatusBadRequest, r.JSON("Invalid request"))
	}

	user := &models.User{}
	if err := tx.Find(user, userReq.ID); err != nil {
		log.Logger.Error().Err(err).Msg("Failed to find user in db")
		return c.Render(http.StatusNotFound, r.JSON("User not found"))
	}

	if userReq.Name != "" {
		user.Name = userReq.Name
	}
	if userReq.Surname != "" {
		user.Surname = userReq.Surname
	}
	if userReq.Patronymic != "" {
		user.Patronymic = userReq.Patronymic
	}
	if userReq.Address != "" {
		user.Address = userReq.Address
	}
	if userReq.PassportSerie != 0 {
		user.PassportSerie = userReq.PassportSerie
	}
	if userReq.PassportNumber != 0 {
		user.PassportNumber = userReq.PassportNumber
	}
	user.UpdatedAt = time.Now()

	if err := tx.Update(user); err != nil {
		log.Logger.Error().Err(err).Msg("Failed to update user in db")
		return c.Render(http.StatusInternalServerError, r.JSON("Failed to update user"))
	}

	return c.Render(http.StatusOK, r.JSON(user))
}

// DeleteUser
//
//	@Summary		Delete a user
//	@Description	Delete an existing user by ID
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			user_id	path		integer	true	"User ID"
//	@Success		200		{string}	string	"User deleted"
//	@Failure		400		{string}	string	"Invalid id"
//	@Failure		404		{string}	string	"User not found"
//	@Router			/user/{id} [delete]
func DeleteUser(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		log.Logger.Error().Msg("Start transaction error")
		return c.Render(http.StatusInternalServerError, r.JSON("Internal server error"))
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Logger.Error().Err(err).Msg("failed to parse user ID")
		return c.Render(http.StatusBadRequest, r.String("Invalid id"))
	}

	err = tx.Destroy(&models.User{ID: id})
	if err != nil {
		log.Logger.Error().Err(err).Msg("Failed to delete user from db")
		return c.Render(http.StatusNotFound, r.String("User not found"))
	}

	return c.Render(http.StatusOK, r.String("User deleted"))
}
