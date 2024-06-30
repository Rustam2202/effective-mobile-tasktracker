package actions

import (
	"tasktracker/models"
)

func (as *ActionSuite) Test_UserCreate() {
	var err error

	as.Session.Set("current_user_id", "1")
	userReq := &createUserRequest{
		PassportNuber: "1234 567890",
	}

	resJSON := as.JSON("/user").Post(userReq)
	as.Equal(200, resJSON.Code)
	err = models.DB.First(&models.User{})
	as.NoError(err)
	count, err := models.DB.Count(&models.User{})
	as.NoError(err)
	as.Equal(1, count)

	res := as.HTML(`/user?name=Иван&page=1&per_page=1`).Get()
	as.Equal(200, res.Code)
	as.Contains(res.Body.String(), "Иван")

	userUpd := &updateUserRequest{
		ID:             22,
		Surname:        "Иванов",
		Name:           "Сергей",
		Patronymic:     "",
		Address:        "",
		PassportSerie:  0,
		PassportNumber: 0,
	}
	res = as.HTML("/user").Put(userUpd)
	as.Equal(200, resJSON.Code)
}

// func (as *ActionSuite) Test_UserUpdate() {
// 	as.Session.Set("current_user_id", "1")
// 	userReq := &updateUserRequest{
// 		ID:             1,
// 		Name:           "User #1",
// 		Surname:        "Surname Useer #1",
// 		Patronymic:     "Patronymic User #1",
// 		PassportSerie:  1234,
// 		PassportNumber: 567890,
// 		Address:        "Address User #1",
// 	}

// 	res := as.JSON("/user/1").Put(userReq)
// 	as.Equal(200, res.Code)

// 	// var user models.User
// 	err := models.DB.Find(&models.User{}, 1)
// 	as.NoError(err)
// }
