package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/zvoleg/task-diary-back/internal/models"
)

type RoleRepository interface {
	Get(context.Context, uuid.UUID) (*models.RoleResponse, error)
	GetList(context.Context) (*models.AllRoleResponse, error)
}
