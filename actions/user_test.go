package actions

import "tasktracker/models"

func (as *ActionSuite) Test_UserCreate() {
	as.Session.Set("current_user_id", "1")
	resp := as.HTML("/info").Get()
	resp.Body.String()
	userReq := &createUserRequest{
		PassportNuber: "1234 567890",
	}

	res := as.JSON("/user").Post(userReq)
	as.Equal(200, res.Code)

	// var user models.User
	err := models.DB.Find(&models.User{}, 1)
	as.NoError(err)
}

func (as *ActionSuite) Test_UserUpdate() {
	as.Session.Set("current_user_id", "1")
	userReq := &updateUserRequest{
		ID:             1,
		Name:           "User #1",
		Surname:        "Surname Useer #1",
		Patronymic:     "Patronymic User #1",
		PassportSerie:  1234,
		PassportNumber: 567890,
		Address:        "Address User #1",
	}

	res := as.JSON("/user/1").Put(userReq)
	as.Equal(200, res.Code)

	// var user models.User
	err := models.DB.Find(&models.User{}, 1)
	as.NoError(err)
}
