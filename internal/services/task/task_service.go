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
func (serv *taskService) Create(taskRequest models.TaskRequest) (*models.TaskResponse, error) {
	ctx := context.Background()
	task, err := serv.repo.Create(ctx, taskRequest)
	if err != nil {
		return nil, err
	}
	return task, err
}

// Delete implements services.TaskService
func (serv *taskService) Delete(taskId uuid.UUID) error {
	ctx := context.Background()
	return serv.repo.Delete(ctx, taskId)
}

// Get implements services.TaskService
func (serv *taskService) Get(taskId uuid.UUID) (*models.TaskResponse, error) {
	ctx := context.Background()
	task, err := serv.repo.Get(ctx, taskId)
	if err != nil {
		return nil, err
	}
	return task, nil
}

// GetList implements services.TaskService
func (serv *taskService) GetList() (*models.AllTaskResponse, error) {
	return nil, nil
}

// Update implements services.TaskService
func (serv *taskService) Update(taskId uuid.UUID, taskRequest models.TaskRequest) (*models.TaskResponse, error) {
	ctx := context.Background()
	task, err := serv.repo.Update(ctx, taskId, taskRequest)
	if err != nil {
		return nil, err
	}
	return task, nil
}
