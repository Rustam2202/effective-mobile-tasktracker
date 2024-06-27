package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
	"tasktracker/models" 
)

// UsersResource is the resource for the User model
type UsersResource struct {
	buffalo.Resource
}

// List gets all Users. This function is mapped to the path
// GET /users
func (v UsersResource) GetAllUsers(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	users := &models.User{}
	if err := tx.All(users); err != nil {
		return err
	}
	return c.Render(http.StatusOK, r.JSON(users))
}

// Show gets the data for one User. This function is mapped to
// the path GET /users/{user_id}
func (v UsersResource) GetUserById(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	user := &models.User{}
	if err := tx.Find(user, c.Param("user_id")); err != nil {
		return err
	}
	return c.Render(http.StatusOK, r.JSON(user))
}

// Create adds a User to the DB. This function is mapped to the
// path POST /users
func (v UsersResource) CreateUser(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	user := &models.User{}

	// Bind user to the incoming request payload
	if err := c.Bind(user); err != nil {
		return err
	}

	// Validate the data from the request
	verrs, err := tx.ValidateAndCreate(user)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		return c.Render(http.StatusUnprocessableEntity, r.JSON(verrs))
	}

	return c.Render(http.StatusCreated, r.JSON(user))
}

// Update changes a User in the DB. This function is mapped to
// the path PUT /users/{user_id}
func (v UsersResource) UpdateUser(c buffalo.Context) error {
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
func (v UsersResource) DeleteUser(c buffalo.Context) error {
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
