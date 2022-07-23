package task

import (
	"context"

	"github.com/google/uuid"
	"github.com/zvoleg/task-diary-back/internal/models"
	"github.com/zvoleg/task-diary-back/internal/repositories"
	"github.com/zvoleg/task-diary-back/internal/services"
)

type taskService struct {
	repo repositories.TaskRepository
}

func NewTaskService(repo repositories.TaskRepository) services.TaskService {
	return &taskService{repo: repo}
}

// Create implements services.TaskService
func (serv *taskService) Create(ctx context.Context, taskRequest models.TaskRequest) (*models.TaskResponse, error) {
	task, err := serv.repo.Create(ctx, taskRequest)
	if err != nil {
		return nil, err
	}
	return task, err
}

// Delete implements services.TaskService
func (serv *taskService) Delete(ctx context.Context, taskId uuid.UUID) error {
	return serv.repo.Delete(ctx, taskId)
}

// Get implements services.TaskService
func (serv *taskService) Get(ctx context.Context, taskId uuid.UUID) (*models.TaskResponse, error) {
	task, err := serv.repo.Get(ctx, taskId)
	if err != nil {
		return nil, err
	}
	return task, nil
}

// GetList implements services.TaskService
func (serv *taskService) GetList(ctx context.Context) (*models.AllTaskResponse, error) {
	return nil, nil
}

// Update implements services.TaskService
func (serv *taskService) Update(ctx context.Context, taskId uuid.UUID, taskRequest models.TaskRequest) (*models.TaskResponse, error) {
	task, err := serv.repo.Update(ctx, taskId, taskRequest)
	if err != nil {
		return nil, err
	}
	return task, nil
}
