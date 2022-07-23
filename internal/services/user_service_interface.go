package services

import (
	"github.com/google/uuid"
	"github.com/zvoleg/task-diary-back/internal/models"
)

type UserService interface {
	Get(uuid.UUID) (*models.UserResponse, error)
	Create(*models.UserRequest) (*models.UserResponse, error)
	Update(uuid.UUID, *models.UserRequest) (*models.UserResponse, error)
	Delete(uuid.UUID) error
}
