package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/zvoleg/task-diary-back/internal/models"
)

type TaskRepository interface {
	Get(context.Context, uuid.UUID) (*models.TaskResponse, error)
	GetList(context.Context) (*models.AllTaskResponse, error)
	Create(context.Context, models.TaskRequest) (*models.TaskResponse, error)
	Update(context.Context, models.TaskRequest) (*models.TaskResponse, error)
	Delete(context.Context, uuid.UUID) error
}
