package actions

import (
	"net/http"
	"strconv"

	"tasktracker/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
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

func CreateUser(c buffalo.Context) error {
}

// Update changes a User in the DB. This function is mapped to
// the path PUT /users/{user_id}
func UpdateUser(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	user := &models.User{}
	if err := tx.Find(user, c.Param("user_id")); err != nil {
		return err
	}

	// Bind user to the incoming request payload
	if err := c.Bind(user); err != nil {
		return err
	}

	verrs, err := tx.ValidateAndUpdate(user)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		return c.Render(http.StatusUnprocessableEntity, r.JSON(verrs))
	}

	return c.Render(http.StatusOK, r.JSON(user))
}

// Destroy deletes a User from the DB. This function is mapped
// to the path DELETE /users/{user_id}
func DeleteUser(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	user := &models.User{}
	if err := tx.Find(user, c.Param("user_id")); err != nil {
		return err
	}

	if err := tx.Destroy(user); err != nil {
		return err
	}

	return c.Render(http.StatusOK, r.JSON(user))
}
