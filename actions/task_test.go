package actions

import (
	"encoding/json"
	"fmt"
	"tasktracker/models"
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
	task_bind := &models.TaskBind{}
	err := as.DB.First(task_bind)

	as.NoError(err)
	as.Equal(200, res.Code)
	as.Equal(1, task_bind.UserID)
	as.Equal(1, task_bind.TaskID)
	as.NotEmpty(task_bind.CreatedAt)
	as.Empty(task_bind.FinishAt)
}

func (as *ActionSuite) Test_StopTaskOfUser() {
	as.Session.Set("current_user_id", "1")
	as.LoadFixture("user_test")
	as.LoadFixture("task_test")

	req := userRequest{
		UserID: 1,
		TaskID: 1,
	}

	as.JSON("/task/start").Post(req)
	time.Sleep(1 * time.Second)

	res := as.JSON("/task/stop").Post(req)
	task_bind := &models.TaskBind{}
	err := as.DB.First(task_bind)

	as.NoError(err)
	as.Equal(200, res.Code)
	as.Equal(1, task_bind.UserID)
	as.Equal(1, task_bind.TaskID)
	as.NotEmpty(task_bind.CreatedAt)
	as.NotEmpty(task_bind.FinishAt)
	as.Greater(task_bind.FinishAt.Unix(), task_bind.CreatedAt.Unix())
}

func (as *ActionSuite) Test_GetTimeUsersTask() {
	as.Session.Set("current_user_id", "1")
	as.LoadFixture("user_test")
	as.LoadFixture("task_test")
	as.LoadFixture("task_binds_test")

	beginPeriod := time.Now().Add(-5 * time.Second).Format(time.DateTime)
	endPeriod := time.Now().Add(-2 * time.Second).Format(time.DateTime)
	query := fmt.Sprintf(`/task?user_id=1&begin_period=%s&end_period=%s`, beginPeriod, endPeriod)
	res := as.HTML(query).Get()

	sums := TaskSums{}
	err := json.Unmarshal(res.Body.Bytes(), &sums)

	as.NoError(err)
	as.Equal(200, res.Code)
	as.Len(sums, 4)
	var totalSum float64
	for _, sum := range sums {
		totalSum += sum.TimeSum
	}
	as.Equal(7.0, totalSum)
}
