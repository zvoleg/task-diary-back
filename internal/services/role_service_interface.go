package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/zvoleg/task-diary-back/internal/models"
)

type RoleService interface {
	Get(context.Context, uuid.UUID) (*models.RoleResponse, error)
	GetList(context.Context) (*models.AllRoleResponse, error)
}
