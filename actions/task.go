package actions

import (
	"net/http"
	"strconv"
	log "tasktracker/logger"
	"tasktracker/models"
	"time"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
)

type userRequest struct {
	UserID int `json:"user_id"`
	TaskID int `json:"task_id"`
}

func StartTaskOfUser(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		log.Logger.Error().Msg("Start transaction error")
		return c.Render(http.StatusInternalServerError, r.JSON("Internal server error"))
	}

	var req userRequest
	err := c.Bind(&req)
	if err != nil {
		log.Logger.Error().Err(err).Msg("failed to bind request")
		return c.Render(400, r.JSON("Invalid request"))
	}

	task := models.TaskBind{
		ID: 1,
		UserID:  req.UserID,
		TaskID:  req.TaskID,
		StartAt: (time.Now().Add(time.Second * 3)),
		FinishAt: (time.Now()),
	}

	err = tx.Update(task)
	if err != nil {
		log.Logger.Error().Err(err).Msg("failed to update task")
		return c.Render(500, r.JSON("Internal server error"))
	}
	return c.Render(200, r.JSON(map[string]string{}))
}

func StopTaskOfUser(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		log.Logger.Error().Msg("Start transaction error")
		return c.Render(http.StatusInternalServerError, r.JSON("Internal server error"))
	}

	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		log.Logger.Error().Err(err).Msg("failed to convert user id to int")
		return c.Render(400, r.JSON("Invalid user id"))
	}
	taskID, err := strconv.Atoi(c.Param("task_id"))
	if err != nil {
		log.Logger.Error().Err(err).Msg("failed to convert task id to int")
		return c.Render(400, r.JSON("Invalid task id"))
	}

	task := models.TaskBind{
		UserID:   userID,
		TaskID:   taskID,
		FinishAt: (time.Now()),
	}

	err = tx.Update(task)
	if err != nil {
		log.Logger.Error().Err(err).Msg("failed to update task")
		return c.Render(500, r.JSON("Internal server error"))
	}
	return c.Render(200, r.JSON(map[string]string{}))
}

func GetTimeUsersTask(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		log.Logger.Error().Msg("Start transaction error")
		return c.Render(http.StatusInternalServerError, r.JSON("Internal server error"))
	}

	beginPeriod, err := time.Parse(time.RFC3339, c.Param("begin_period"))
	if err != nil {
		log.Logger.Error().Err(err).Msg("failed to parse begin period")
		return c.Render(400, r.JSON("Invalid begin period"))
	}
	endPeriod, err := time.Parse(time.RFC3339, c.Param("end_period"))
	if err != nil {
		log.Logger.Error().Err(err).Msg("failed to parse end period")
		return c.Render(400, r.JSON("Invalid end period"))
	}

	var result time.Time
	err = tx.RawQuery(
		`SELECT SUM(finish_at - start_at) FROM task_binds WHERE start_at >= ? AND finish_at <= ?`,
		beginPeriod, endPeriod, beginPeriod, endPeriod).First(&result)
	if err != nil {
		log.Logger.Error().Err(err).Msg("failed to get time users task")
		return c.Render(http.StatusNotFound, r.JSON("Not found"))
	}
	return c.Render(200, r.JSON(map[string]interface{}{
		"hour":   result.Hour(),
		"minute": result.Minute(),
	}))
}
