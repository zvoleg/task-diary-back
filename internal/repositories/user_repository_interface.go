package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/zvoleg/task-diary-back/internal/models"
)

type UserRepository interface {
	Create(context.Context, *models.UserRequest) (*models.UserResponse, error)
	Update(context.Context, uuid.UUID, *models.UserRequest) (*models.UserResponse, error)
	Get(context.Context, uuid.UUID) (*models.UserResponse, error)
	GetList(context.Context) ([]*models.AllUserResponse, error)
	Delete(context.Context, uuid.UUID) error
}
