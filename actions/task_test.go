package actions

import (
	"fmt"
	"net/http"
	"time"
)

func (as *ActionSuite) Test_StartTaskOfUser() {
	as.Session.Set("current_user_id", "1")
	as.LoadFixture("user_test")
	as.LoadFixture("task_test")

	req := userRequest{
		UserID: 1,
		TaskID: 1,
	}

	res := as.JSON("/task/start").Post(req)
	as.Equal(200, res.Code)
}

func (as *ActionSuite) Test_StopTaskOfUser() {
	as.Session.Set("current_user_id", "1")
	as.LoadFixture("user_test")
	as.LoadFixture("task_test")
	as.LoadFixture("task_binds_test")

	req := userRequest{
		UserID: 1,
		TaskID: 1,
	}

	res := as.JSON("/task/stop").Post(req)
	as.Equal(200, res.Code)

}

func (as *ActionSuite) Test_GetTimeUsersTask() {
	as.Session.Set("current_user_id", "1")
	as.LoadFixture("user_test")
	as.LoadFixture("task_test")
	as.LoadFixture("task_binds_test")

	beginPeriod := time.Now().Add(-5 * time.Second).Format(http.TimeFormat)
	endPeriod := time.Now().Add(-2 * time.Second).Format(http.TimeFormat)
	query := fmt.Sprintf(`/task?user_id=1&begin_period=%s&end_period=%s`, beginPeriod, endPeriod)
	res := as.HTML(query).Get()
	as.Equal(200, res.Code)
}

// time.ParseError {Layout: "2006-01-02T15:04:05Z07:00", Value: "2024-07-01T14:56:31 04:00", LayoutElem: "Z07:00", ValueElem: " 04:00", Message: ""}
// "/task?user_id=1&begin_period=                                2024-07-01T15:11:25+04:00   &end_period=2024-07-01T15:11:28+04:00"