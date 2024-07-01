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

func StartUserTask(c buffalo.Context) error {
	var (
		err  error
		user models.User
		task models.Task
		bind models.TaskBind
	)
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		log.Logger.Error().Msg("Start transaction error")
		return c.Render(http.StatusInternalServerError, r.JSON("Internal server error"))
	}

	var req userRequest
	err = c.Bind(&req)
	if err != nil {
		log.Logger.Error().Err(err).Msg("failed to bind request")
		return c.Render(400, r.JSON("Invalid request"))
	}

	err = tx.Find(&user, req.UserID)
	if err != nil {
		log.Logger.Error().Err(err).Msg("failed to find user")
		return c.Render(404, r.JSON("User not found"))
	}
	err = tx.Find(&task, req.TaskID)
	if err != nil {
		log.Logger.Error().Err(err).Msg("failed to find task")
		return c.Render(404, r.JSON("Task not found"))
	}

	bind = models.TaskBind{
		ID:        0,
		TaskID:    req.TaskID,
		UserID:    req.UserID,
		StartAt:   time.Now(),
		FinishAt:  time.Time{},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = tx.Create(&bind)
	// err=models.DB.Create(&bind)
	if err != nil {
		log.Logger.Error().Err(err).Msg("failed to update task")
		return c.Render(500, r.JSON("Internal server error"))
	}
	return c.Render(200, r.JSON(bind))
}

func StopUserTask(c buffalo.Context) error {
	var (
		err  error
		user models.User
		task models.Task
		bind models.TaskBind
	)

	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		log.Logger.Error().Msg("Start transaction error")
		return c.Render(http.StatusInternalServerError, r.JSON("Internal server error"))
	}

	var req userRequest
	err = c.Bind(&req)
	if err != nil {
		log.Logger.Error().Err(err).Msg("failed to bind request")
		return c.Render(400, r.JSON("Invalid request"))
	}

	err = tx.Find(&user, req.UserID)
	if err != nil {
		log.Logger.Error().Err(err).Msg("failed to find user")
		return c.Render(404, r.JSON("User not found"))
	}
	err = tx.Find(&task, req.TaskID)
	if err != nil {
		log.Logger.Error().Err(err).Msg("failed to find task")
		return c.Render(404, r.JSON("Task not found"))
	}
	bind.UserID = user.ID
	bind.TaskID = task.ID
	err = tx.First(&bind)
	if err != nil {
		log.Logger.Error().Err(err).Msg("failed to find task bind")
		return c.Render(404, r.JSON("Task bind not found"))
	}
	bind.FinishAt = time.Now()

	err = tx.Update(&bind)
	if err != nil {
		log.Logger.Error().Err(err).Msg("failed to update task")
		return c.Render(500, r.JSON("Internal server error"))
	}
	return c.Render(200, r.JSON(map[string]string{}))
}

// type timeUsersTask struct {
// 	UserID      int       `json:"user_id"`
// 	BeginPeriod time.Time `json:"begin_period"`
// 	EndPeriod   time.Time `json:"end_period"`
// }

func GetTimeUsersTask(c buffalo.Context) error {
	var (
		err         error
		userId      int
		beginPeriod time.Time
		endPeriod   time.Time
	)

	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		log.Logger.Error().Msg("Start transaction error")
		return c.Render(http.StatusInternalServerError, r.JSON("Internal server error"))
	}

	userId, err = strconv.Atoi(c.Param("user_id"))
	if err != nil {
		log.Logger.Error().Err(err).Msg("failed to convert user_id")
		return c.Render(400, r.JSON("Invalid request"))
	}
	beginPeriod, err = time.Parse(http.TimeFormat, c.Param("begin_period"))
	log.Logger.Debug().Msgf("begin_period: %s", c.Param("begin_period"))
	if err != nil {
		log.Logger.Error().Err(err).Msg("failed to convert begin_period")
		return c.Render(400, r.JSON("Invalid request"))
	}
	endPeriod, err = time.Parse(http.TimeFormat, c.Param("end_period"))
	log.Logger.Debug().Msgf("end_period: %s", c.Param("end_period"))
	if err != nil {
		log.Logger.Error().Err(err).Msg("failed to convert end_period")
		return c.Render(400, r.JSON("Invalid request"))
	}

	query := `
    WITH adjusted_times AS (
        SELECT 
            GREATEST(start_at, ?) AS begin_time,
            LEAST(CASE WHEN finish_at IS NULL THEN NOW() ELSE finish_at END, ?) AS end_time
        FROM 
            task_binds
        WHERE 
            user_id = ? AND
            (start_at <= ? AND (finish_at >= ? OR finish_at IS NULL))
    )
    SELECT 
        SUM(EXTRACT(EPOCH FROM (end_time - begin_time))) AS total_seconds
    FROM 
        adjusted_times
    WHERE 
        begin_time < end_time;
    `
	var totalSeconds float64
	err = tx.RawQuery(query, beginPeriod, endPeriod, userId, endPeriod, beginPeriod).First(&totalSeconds)

	if err != nil {
		log.Logger.Error().Err(err).Msg("failed to calculate total time")
		return c.Render(500, r.JSON("Internal server error"))
	}

	hours := int(totalSeconds / 3600)
	minutes := int(totalSeconds/60) % 60

	return c.Render(200, r.JSON(map[string]interface{}{
		"hours":   hours,
		"minutes": minutes,
	}))
}
