package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/zvoleg/task-diary-back/internal/models"
)

type UserService interface {
	Get(context.Context, uuid.UUID) (*models.UserResponse, error)
	Create(context.Context, *models.UserRequest) (*models.UserResponse, error)
	Update(context.Context, uuid.UUID, *models.UserRequest) (*models.UserResponse, error)
	Delete(context.Context, uuid.UUID) error
}
