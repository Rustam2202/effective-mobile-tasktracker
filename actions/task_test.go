package actions

func (as *ActionSuite) Test_StartTaskOfUser() {
	as.Session.Set("current_user_id", "1")
	as.LoadFixture("user_test")
	as.LoadFixture("task_test")

	req := userRequest{UserID: 1, TaskID: 1}

	res := as.JSON("/task/start").Post(req)
	as.Equal(200, res.Code)
}

func (as *ActionSuite) Test_StopTaskOfUser() {
}

func (as *ActionSuite) Test_GetTimeUsersTask() {
}
