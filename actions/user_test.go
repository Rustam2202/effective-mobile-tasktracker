package actions

import "tasktracker/models"

func (as *ActionSuite) Test_UserCreate() {
	as.Session.Set("current_user_id", "1")
	userReq := &createUserRequest{
		Name:          "User #1",
		Surname:       "Surname Useer #1",
		Patronymic:    "Patronymic User #1",
		PassportNuber: "1234 567890",
		Address:       "Address User #1",
	}

	res := as.JSON("/user").Post(userReq)
	as.Equal(200, res.Code)

	var user models.User
	err:=as.DB.First(&user)
	as.NoError(err)
}
