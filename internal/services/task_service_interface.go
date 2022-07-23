package services

import (
	"github.com/google/uuid"
	"github.com/zvoleg/task-diary-back/internal/models"
)

type TaskService interface {
	Get(uuid.UUID) (*models.TaskResponse, error)
	GetList() (*models.AllTaskResponse, error)
	Create(models.TaskRequest) (*models.TaskResponse, error)
	Update(uuid.UUID, models.TaskRequest) (*models.TaskResponse, error)
	Delete(uuid.UUID) error
}
