package models

import (
	"time"
)

type TaskBind struct {
	ID       int       `json:"id" db:"id"`
	TaskID   int       `json:"task_id" db:"task_id"`
	UserID   int       `json:"user_id" db:"user_id"`
	StartAt  time.Time `json:"start_at" db:"start_at"`
	FinishAt time.Time `json:"finish_at" db:"finish_at"`
}
