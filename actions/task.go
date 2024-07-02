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

// StartUserTask
//
//	@Summary		Start a task for a user
//	@Description	Start a task for a user
//	@Tags			Tasks
//	@Accept			json
//	@Produce		json
//	@Param			userReq	body		userRequest	true	"User and Task IDs"
//	@Success		200		{object}	models.TaskBind
//	@Failure		400		{string}	string	"Invalid request"
//	@Failure		404		{string}	string	"User or Task not found"
//	@Failure		500		{string}	string	"Internal server error"
//	@Router			/task/start [post]
func StartUserTask(c buffalo.Context) error {
	var (
		err      error
		user     models.User
		task     models.Task
		taskBind models.TaskBind
	)

	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		log.Logger.Error().Msg("Failed to start transaction")
		return c.Render(http.StatusInternalServerError, r.String("Internal server error"))
	}

	var req userRequest
	err = c.Bind(&req)
	if err != nil {
		log.Logger.Error().Err(err).Msg("Failed to bind request")
		return c.Render(400, r.String("Invalid request"))
	}

	err = tx.Find(&user, req.UserID)
	if err != nil {
		log.Logger.Error().Err(err).Msg("Failed to find user")
		return c.Render(404, r.String("User not found"))
	}
	err = tx.Find(&task, req.TaskID)
	if err != nil {
		log.Logger.Error().Err(err).Msg("failed to find task")
		return c.Render(404, r.String("Task not found"))
	}

	taskBind = models.TaskBind{
		ID:        0,
		TaskID:    req.TaskID,
		UserID:    req.UserID,
		StartAt:   time.Now(),
		FinishAt:  time.Time{},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = tx.Create(&taskBind)
	if err != nil {
		log.Logger.Error().Err(err).Msg("failed to update task")
		return c.Render(500, r.JSON("Internal server error"))
	}
	log.Logger.Debug().Msg("Task Bind was created (User started a Task)")

	return c.Render(200, r.JSON(taskBind))
}

// StopUserTask
//
//	@Summary		Stop a task for a user
//	@Description	Stop a task for a user
//	@Tags			Tasks
//	@Accept			json
//	@Produce		json
//	@Param			userReq	body		userRequest	true	"User and Task IDs"
//	@Success		200		{object}	models.TaskBind
//	@Failure		400		{string}	string	"Invalid request"
//	@Failure		404		{string}	string	"User or Task not found"
//	@Failure		500		{string}	string	"Internal server error"
//	@Router			/task/stop [post]
func StopUserTask(c buffalo.Context) error {
	var (
		err  error
		user models.User
		task models.Task
		bind models.TaskBind
	)

	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		log.Logger.Error().Msg("Failed to start transaction")
		return c.Render(http.StatusInternalServerError, r.String("Internal server error"))
	}

	var req userRequest
	err = c.Bind(&req)
	if err != nil {
		log.Logger.Error().Err(err).Msg("Failed to bind request")
		return c.Render(400, r.String("Invalid request"))
	}

	err = tx.Find(&user, req.UserID)
	if err != nil {
		log.Logger.Error().Err(err).Msg("Failed to find user")
		return c.Render(404, r.String("User not found"))
	}
	err = tx.Find(&task, req.TaskID)
	if err != nil {
		log.Logger.Error().Err(err).Msg("Failed to find task")
		return c.Render(404, r.String("Task not found"))
	}

	bind.UserID = user.ID
	bind.TaskID = task.ID
	err = tx.First(&bind)
	if err != nil {
		log.Logger.Error().Err(err).Msg("failed to find task bind")
		return c.Render(404, r.String("Task bind not found"))
	}
	bind.FinishAt = time.Now()

	err = tx.Update(&bind)
	if err != nil {
		log.Logger.Error().Err(err).Msg("failed to update task")
		return c.Render(500, r.String("Internal server error"))
	}
	log.Logger.Debug().Msg("Task Bind was updated (User stopped a Task)")

	return c.Render(200, r.JSON(map[string]string{}))
}

type TaskSums []struct {
	TaskID  int     `json:"task_id" db:"task_id"`
	TimeSum float64 `json:"time_sum" db:"time_sum"`
}

// GetTimeUsersTask
//
//	@Summary		Get total time spent by a user on tasks within a period
//	@Description	Get total time spent by a user on tasks within a period
//	@Tags			Tasks
//	@Accept			json
//	@Produce		json
//	@Param			user_id			path		int		true	"User ID"
//	@Param			begin_period	query		string	true	"Begin period"
//	@Param			end_period		query		string	true	"End period"
//	@Success		200				{object}	TaskSums
//	@Failure		400				{string}	string	"Invalid request"
//	@Failure		500				{string}	string	"Internal server error"
//	@Router			/task [get]
func GetTimeUsersTask(c buffalo.Context) error {
	var (
		err         error
		userId      int
		beginPeriod time.Time
		endPeriod   time.Time
	)

	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		log.Logger.Error().Msg("Failed to start transaction")
		return c.Render(http.StatusInternalServerError, r.String("Internal server error"))
	}

	userId, err = strconv.Atoi(c.Param("user_id"))
	if err != nil {
		log.Logger.Error().Err(err).Msg("failed to convert user_id")
		return c.Render(400, r.String("Invalid request"))
	}
	beginPeriod, err = time.Parse(http.TimeFormat, c.Param("begin_period"))
	log.Logger.Debug().Msgf("begin_period: %s", c.Param("begin_period"))
	if err != nil {
		log.Logger.Error().Err(err).Msg("failed to convert begin_period")
		return c.Render(400, r.String("Invalid request"))
	}
	endPeriod, err = time.Parse(http.TimeFormat, c.Param("end_period"))
	log.Logger.Debug().Msgf("end_period: %s", c.Param("end_period"))
	if err != nil {
		log.Logger.Error().Err(err).Msg("failed to convert end_period")
		return c.Render(400, r.String("Invalid request"))
	}

	query := `
	 WITH adjusted_times AS (
            SELECT 
                GREATEST(start_at, ?) AS begin_time,
                LEAST(CASE WHEN finish_at IS NULL THEN ? ELSE finish_at END, ?) AS end_time,
                task_id
            FROM 
                task_binds
            WHERE 
                user_id = ? AND
                (start_at <= ? AND (finish_at >= ? OR finish_at IS NULL))
        )
        SELECT 
            task_id,
            SUM(EXTRACT(EPOCH FROM (end_time - begin_time))) AS time_sum
        FROM 
            adjusted_times
        WHERE 
            begin_time < end_time
        GROUP BY 
            task_id
        ORDER BY 
            time_sum DESC;
	`

	var taskSums TaskSums
	err = tx.RawQuery(query, beginPeriod, endPeriod, endPeriod, userId, endPeriod, beginPeriod).All(&taskSums)
	if err != nil {
		log.Logger.Error().Err(err).Msg("failed to calculate total time")
		return c.Render(500, r.String("Internal server error"))
	}
	log.Logger.Debug().Msg("Total time was calculated")

	return c.Render(200, r.JSON(taskSums))
}
