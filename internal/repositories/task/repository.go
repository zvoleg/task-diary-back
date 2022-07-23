package task

import (
	"context"
	"database/sql"
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
	taskDb, err := rep.getTaskDb(ctx, identifier)
	if err != nil {
		return nil, errors.Wrap(err, "repo.Task.Get")
	}
	if taskDb.IsDeleted {
		return nil, errors.Wrap(models.NewErrNotFoundInDb(identifier), "repo.Task.Get")
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
		uuid.New(), // task_id
		uuid.New(), // board_id should be from context
		taskRequest.Title,
		taskRequest.Description,
		taskRequest.Type,
		taskRequest.Status,
		uuid.New(),       // author_id should be from context
		time.Now().UTC(), // created_at
		nil,              // updated_at
		taskRequest.Tags,
		false, // is_deleted
	).StructScan(taskDb); err != nil {
		return nil, errors.Wrap(err, "repo.Task.Create: struct scan error")
	}
	task := taskDb.Map()
	return &task, nil
}

// Update implements repositories.TaskRepository
func (rep *taskRepository) Update(ctx context.Context, identifier uuid.UUID, taskRequest models.TaskRequest) (*models.TaskResponse, error) {
	taskDb, err := rep.getTaskDb(ctx, identifier)
	if err != nil {
		return nil, errors.Wrap(err, "repo.Task.Update")
	}
	if taskDb.IsDeleted {
		return nil, errors.Wrap(models.NewErrNotFoundInDb(identifier), "repo.Task.Update")
	}
	if err := rep.db.QueryRowxContext(
		ctx,
		updateScript,
		taskRequest.Title,
		taskRequest.Description,
		taskRequest.Type,
		taskRequest.Status,
		time.Now().UTC(), // updated_at
		taskRequest.Tags,
	).StructScan(taskDb); err != nil {
		return nil, errors.Wrap(err, "repo.Task.Update: struct scan error")
	}
	task := taskDb.Map()
	return &task, nil
}

// Delete implements repositories.TaskRepository
func (rep *taskRepository) Delete(ctx context.Context, identifier uuid.UUID) error {
	_, err := rep.getTaskDb(ctx, identifier)
	if err != nil {
		return errors.Wrap(err, "repo.Task.Delete")
	}
	if _, err := rep.db.QueryxContext(
		ctx,
		deleteScript,
		identifier,
		time.Now().UTC(),
	); err != nil {
		return errors.Wrap(err, "repo.Task.Delete")
	}
	return nil
}

func (rep *taskRepository) getTaskDb(ctx context.Context, identifier uuid.UUID) (*models.TaskDb, error) {
	taskDb := new(models.TaskDb)
	if err := rep.db.QueryRowxContext(
		ctx,
		getScript,
		identifier,
	).StructScan(taskDb); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, models.NewErrNotFoundInDb(identifier)
		default:
			return nil, err
		}
	}
	return taskDb, nil
}
