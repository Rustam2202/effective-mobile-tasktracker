package actions

import (
	"encoding/json"
	"strconv"
	"tasktracker/models"
)

func (as *ActionSuite) Test_UserCRUDFromInfo() {
	var (
		err   error
		user  models.User
		users []models.User
	)
	as.Session.Set("current_user_id", "1")
	userReq := &createUserRequest{
		PassportNuber: "1234 567890",
	}

	// create user
	resCreate := as.JSON("/user").Post(userReq)
	as.Equal(200, resCreate.Code)
	err = models.DB.First(&models.User{})
	as.NoError(err)
	count, err := models.DB.Count(&models.User{})
	as.NoError(err)
	as.Equal(1, count)

	// get user
	resGet := as.HTML(`/user?name=Иван&page=1&per_page=1`).Get()
	as.Equal(200, resGet.Code)
	err = json.Unmarshal(resGet.Body.Bytes(), &users)
	as.NoError(err)
	as.Len(users, 1)
	as.Equal("Иван", users[0].Name)
	as.Equal("Иванов", users[0].Surname)
	as.Equal("Иванович", users[0].Patronymic)
	as.Equal("г. Москва, ул. Ленина, д. 5, кв. 1", users[0].Address)
	as.Equal(1234, users[0].PassportSerie)
	as.Equal(567890, users[0].PassportNumber)
	as.NotZero(users[0].CreatedAt)
	as.NotZero(users[0].UpdatedAt)

	// update user
	updatedUser := updateUserRequest{
		Surname:        "Петров",
		Name:           "Петр",
		Patronymic:     "Петрович",
		PassportSerie:  4321,
		PassportNumber: 987654,
		Address:        "г. Санкт-Петербург, ул. Невская, д. 10, кв. 2",
	}

	resUpdate := as.JSON("/user/" + strconv.Itoa(users[0].ID)).Put(updatedUser)
	as.Equal(200, resUpdate.Code)
	err = json.Unmarshal(resUpdate.Body.Bytes(), &user)
	as.NoError(err)
	as.Equal("Петр", user.Name)
	as.Equal("Петров", user.Surname)
	as.Equal("Петрович", user.Patronymic)
	as.Equal("г. Санкт-Петербург, ул. Невская, д. 10, кв. 2", user.Address)
	as.Equal(4321, user.PassportSerie)
	as.Equal(987654, user.PassportNumber)
	as.NotZero(user.CreatedAt)
	as.NotZero(user.UpdatedAt)

	// delete user
	resDelete := as.JSON("/user/" + strconv.Itoa(user.ID)).Delete()
	as.Equal(200, resDelete.Code)
	err = as.DB.Find(&models.User{}, user.ID)
	as.ErrorContains(err, "sql: no rows in result set")
}
