package models

import (
	"time"
)

type Todo struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Task      string    `json:"task"`
	CreatedAt time.Time `json:"created_at"`
	Complete  bool      `json:"completed"`
}

type CreateTodoInput struct {
	Task     string `json:"task" binding:"required"`
	Complete bool   `json:"completed"`
}

type UpdateTodoInput struct {
	Task     string `json:"task"`
	Complete bool   `json:"completed"`
}
