package role

import (
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/zvoleg/task-diary-back/internal/models"
	"github.com/zvoleg/task-diary-back/internal/repositories"
	"github.com/zvoleg/task-diary-back/internal/services"
)

type roleService struct {
	repo repositories.RoleRepository
}

func NewRoleService(repo repositories.RoleRepository) services.RoleService {
	return &roleService{repo: repo}
}

func (serv *roleService) Get(ctx context.Context, roleId uuid.UUID) (*models.RoleResponse, error) {
	role, err := serv.repo.Get(ctx, roleId)
	if err != nil {
		return nil, errors.Wrap(err, "roleServ.Get: can't find role")
	}
	return role, nil
}

func (serv *roleService) GetList(ctx context.Context) (*models.AllRoleResponse, error) {
	allRoles, err := serv.repo.GetList(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "roleServ.GetList: can't get role list")
	}
	return allRoles, err
}
