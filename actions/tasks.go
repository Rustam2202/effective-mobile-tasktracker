package actions

import (
	"net/http"
	"strconv"
	"tasktracker/models"
	"time"

	"github.com/gobuffalo/buffalo"
)

func StartTaskOfUser(c buffalo.Context) error {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		return c.Render(400, r.JSON("Invalid user id"))
	}
	taskID, err := strconv.Atoi(c.Param("task_id"))
	if err != nil {
		return c.Render(400, r.JSON("Invalid task id"))
	}

	task := models.TaskBind{
		UserID:  userID,
		TaskID:  taskID,
		StartAt: time.Now(),
	}

	err = models.DB.Update(task)
	if err != nil {
		return c.Render(500, r.JSON("Internal server error"))
	}
	return c.Render(200, r.JSON(map[string]string{}))
}

func StopTaskOfUser(c buffalo.Context) error {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		return c.Render(400, r.JSON("Invalid user id"))
	}
	taskID, err := strconv.Atoi(c.Param("task_id"))
	if err != nil {
		return c.Render(400, r.JSON("Invalid task id"))
	}

	task := models.TaskBind{
		UserID:   userID,
		TaskID:   taskID,
		FinishAt: time.Now(),
	}

	err = models.DB.Update(task)
	if err != nil {
		return c.Render(500, r.JSON("Internal server error"))
	}
	return c.Render(200, r.JSON(map[string]string{}))
}

func GetTimeUsersTask(c buffalo.Context) error {
	beginPeriod, err := time.Parse(time.RFC3339, c.Param("begin_period"))
	if err != nil {
		return c.Render(400, r.JSON("Invalid begin period"))
	}
	endPeriod, err := time.Parse(time.RFC3339, c.Param("end_period"))
	if err != nil {
		return c.Render(400, r.JSON("Invalid end period"))
	}

	var result time.Time
	err = models.DB.RawQuery(
		`SELECT SUM(finish_at - start_at) FROM task_binds WHERE start_at >= ? AND finish_at <= ?`,
		beginPeriod, endPeriod, beginPeriod, endPeriod).First(&result)
	if err != nil {
		return c.Render(http.StatusNotFound, r.JSON("Not found"))
	}
	return c.Render(200, r.JSON(map[string]interface{}{
		"hour":   result.Hour(),
		"minute": result.Minute(),
	}))
}
