package services

import (
	"github.com/google/uuid"
	"github.com/zvoleg/task-diary-back/internal/models"
)

type RoleService interface {
	Get(uuid.UUID) (*models.RoleResponse, error)
	GetList() (*models.AllRoleResponse, error)
}
