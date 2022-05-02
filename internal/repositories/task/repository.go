package task

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/zvoleg/task-diary-back/internal/models"
	"github.com/zvoleg/task-diary-back/internal/repositories"
)

type taskRepository struct {
	db *sqlx.DB
}

func NewTaskRepository(db *sqlx.DB) repositories.TaskRepository {
	return &taskRepository{db: db}
}

// Get implements repositories.TaskRepository
func (rep *taskRepository) Get(ctx context.Context, identifier uuid.UUID) (*models.TaskResponse, error) {
	taskDb := new(models.TaskDb)
	if err := rep.db.QueryRowxContext(
		ctx,
		getScript,
		identifier,
	).StructScan(taskDb); err != nil {
		return nil, errors.Wrap(err, "taskRepo.Get: struct scan error")
	}
	task := taskDb.Map()
	return &task, nil
}

// GetList implements repositories.TaskRepository
func (rep *taskRepository) GetList(ctx context.Context) (*models.AllTaskResponse, error) {
	return nil, nil
}

// Create implements repositories.TaskRepository
func (rep *taskRepository) Create(ctx context.Context, taskRequest models.TaskRequest) (*models.TaskResponse, error) {
	taskDb := new(models.TaskDb)
	if err := rep.db.QueryRowxContext(
		ctx,
		createScript,
		uuid.New(),
		uuid.New(),
		taskRequest.Title,
		taskRequest.Description,
		taskRequest.Type,
		taskRequest.Status,
		uuid.New(),
		time.Now().UTC(),
		nil,
		taskRequest.Tags,
		false,
	).StructScan(taskDb); err != nil {
		return nil, errors.Wrap(err, "taskRepo.Create: struct scan error")
	}
	task := taskDb.Map()
	return &task, nil
}

// Update implements repositories.TaskRepository
func (rep *taskRepository) Update(ctx context.Context, taskRequest models.TaskRequest) (*models.TaskResponse, error) {
	taskDb := new(models.TaskDb)
	if err := rep.db.QueryRowxContext(
		ctx,
		updateScript,
		uuid.New(), // task_id
		uuid.New(), // board_id should be from context
		taskRequest.Title,
		taskRequest.Description,
		taskRequest.Type,
		taskRequest.Status,
		uuid.New(),       // author_id should be from context
		time.Now().UTC(), // created_at
		nil,              //modified_at
		taskRequest.Tags,
		false, // is_deleted
	).StructScan(taskDb); err != nil {
		return nil, errors.Wrap(err, "taskRepo.Update: struct scan error")
	}
	task := taskDb.Map()
	return &task, nil
}

// Delete implements repositories.TaskRepository
func (rep *taskRepository) Delete(ctx context.Context, identifier uuid.UUID) error {
	panic("unimplemented")
}
