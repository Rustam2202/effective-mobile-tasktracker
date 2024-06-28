package actions

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"tasktracker/models"

	"github.com/gobuffalo/buffalo"
)

func GetAllUsers(c buffalo.Context) error {
	users := &models.User{}
	var params []string

	if c.Param("id") != "" {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.Render(http.StatusBadRequest, r.JSON("Invalid id"))
		}
		if id < 0 {
			return c.Render(http.StatusBadRequest, r.JSON("Invalid id"))
		}
		params = append(params, c.Param("id"))
	}
	if c.Param("name") != "" {
		params = append(params, c.Param("name"))
	}
	if c.Param("surname") != "" {
		params = append(params, c.Param("surname"))
	}
	if c.Param("patronymic") != "" {
		params = append(params, c.Param("patronymic"))
	}
	if c.Param("passportSerie") != "" {
		passSer, err := strconv.Atoi(c.Param("passportSerie"))
		if err != nil {
			return c.Render(http.StatusBadRequest, r.JSON("Invalid passport serie"))
		}
		if passSer < 0 || passSer > 4 {
			return c.Render(http.StatusBadRequest, r.JSON("Invalid passport serie"))
		}
		params = append(params, c.Param("passportSerie"))
	}
	if c.Param("passportNumber") != "" {
		passNumb, err := strconv.Atoi(c.Param("passportNumber"))
		if err != nil {
			return c.Render(http.StatusBadRequest, r.JSON("Invalid passport number"))
		}
		if passNumb < 0 || passNumb > 6 {
			return c.Render(http.StatusBadRequest, r.JSON("Invalid passport number"))
		}
		params = append(params, c.Param("passportNumber"))
	}
	if c.Param("address") != "" {
		params = append(params, c.Param("address"))
	}
	if c.Param("page") == "" {
		return c.Render(http.StatusBadRequest, r.JSON("Invalid page number"))
	}
	page, err := strconv.Atoi(c.Param("page"))
	if err != nil {
		return c.Render(http.StatusBadRequest, r.JSON("Invalid page number"))
	}
	if page < 0 {
		return c.Render(http.StatusBadRequest, r.JSON("Invalid page number"))
	}
	if c.Param("perPage") == "" {
		return c.Render(http.StatusBadRequest, r.JSON("Invalid perPage number"))
	}
	perPage, err := strconv.Atoi(c.Param("perPage"))
	if err != nil {
		return c.Render(http.StatusBadRequest, r.JSON("Invalid perPage number"))
	}
	if perPage < 0 {
		return c.Render(http.StatusBadRequest, r.JSON("Invalid perPage number"))
	}

	err = models.DB.Select(params...).Paginate(page, perPage).All(&users)
	if err != nil {
		return c.Render(http.StatusNotFound, r.JSON("Users not found"))
	}
	return c.Render(http.StatusOK, r.JSON(users))
}

func GetUserByPassport(c buffalo.Context) error {
	var err error

	passSerParam := c.Param("passportSerie")
	passSer, err := strconv.Atoi(passSerParam)
	if err != nil {
		return c.Render(http.StatusBadRequest, r.JSON("Invalid passport serie"))
	}
	if passSer < 0 || passSer > 4 {
		return c.Render(http.StatusBadRequest, r.JSON("Invalid passport serie"))
	}

	passNumbParam := c.Param("passportNumber")
	passNumb, err := strconv.Atoi(passNumbParam)
	if err != nil {
		return c.Render(http.StatusBadRequest, r.JSON("Invalid passport number"))
	}
	if passNumb < 0 || passNumb > 6 {
		return c.Render(http.StatusBadRequest, r.JSON("Invalid passport number"))
	}

	user := &models.User{}
	err = models.DB.Where("passport_serie = ? AND passport_number = ?", passSer, passNumb).First(&user)
	if err != nil {
		return c.Render(http.StatusNotFound, r.JSON("User not found"))
	}
	return c.Render(http.StatusOK, r.JSON(user))
}

type createUserRequest struct {
	Name          string `json:"name"`
	Surname       string `json:"surname"`
	Patronymic    string `json:"patronymic"`
	PassportNuber string `json:"passportNumber"`
	Address       string `json:"address"`
}

func CreateUser(c buffalo.Context) error {
	var userReq createUserRequest
	err := c.Bind(&userReq)
	if err != nil {
		return c.Render(http.StatusBadRequest, r.JSON("Invalid request"))
	}

	splitNumber := strings.Split(userReq.PassportNuber, " ")
	if len(splitNumber) != 2 {
		return c.Render(http.StatusBadRequest, r.JSON("Invalid passport number"))
	}
	passSer, err := strconv.Atoi(splitNumber[0])
	if err != nil {
		return c.Render(http.StatusBadRequest, r.JSON("Invalid passport serie"))
	}
	passNumb, err := strconv.Atoi(splitNumber[1])
	if err != nil {
		return c.Render(http.StatusBadRequest, r.JSON("Invalid passport number"))
	}

	user := models.User{
		Name:           userReq.Name,
		Surname:        userReq.Surname,
		Patronymic:     userReq.Patronymic,
		PassportSerie:  passSer,
		PassportNumber: passNumb,
		Address:        userReq.Address,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	err = models.DB.Create(&user)
	if err != nil {
		return c.Render(http.StatusInternalServerError, r.JSON("Internal server error"))
	}

	return c.Render(http.StatusOK, r.JSON(user))
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

func UpdateUser(c buffalo.Context) error {
	var userReq updateUserRequest
	err := c.Bind(&userReq)
	if err != nil {
		return c.Render(http.StatusBadRequest, r.JSON("Invalid request"))
	}

	user := &models.User{}
	if err := models.DB.Find(user, c.Param("user_id")); err != nil {
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

func DeleteUser(c buffalo.Context) error {
	id, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		return c.Render(http.StatusBadRequest, r.JSON("Invalid id"))
	}
	err = models.DB.Destroy(&models.User{ID: id})
	if err != nil {
		return c.Render(http.StatusNotFound, r.JSON("User not found"))
	}

	return c.Render(http.StatusOK, r.JSON("User deleted"))
}
