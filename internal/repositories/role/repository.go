package role

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/zvoleg/task-diary-back/internal/models"
	"github.com/zvoleg/task-diary-back/internal/repositories"
)

type roleRepository struct {
	db *sqlx.DB
}

func NewRoleRepository(db *sqlx.DB) repositories.RoleRepository {
	return &roleRepository{db: db}
}

func (rep *roleRepository) Get(ctx context.Context, identifier uuid.UUID) (*models.RoleResponse, error) {
	role := new(models.RoleResponse)

	if err := rep.db.QueryRowxContext(
		ctx,
		getScript,
		identifier,
	).StructScan(role); err != nil {
		return nil, errors.Wrap(err, "roleRepo.Get: struct scan error")
	}
	return role, nil
}

func (rep *roleRepository) GetList(ctx context.Context) (*models.AllRoleResponse, error) {
	rows, err := rep.db.QueryxContext(
		ctx,
		getListScript,
	)

	if err != nil {
		return nil, errors.Wrap(err, "roleRepo.GetList: query error")
	}

	roleList := make([]*models.RoleResponse, 0)

	for rows.Next() {
		role := new(models.RoleResponse)
		if err := rows.StructScan(role); err != nil {
			return nil, errors.Wrap(err, "roleRepo.GetList: struct scan error")
		}
		roleList = append(roleList, role)
	}

	allRoleResponse := new(models.AllRoleResponse)
	allRoleResponse.Roles = roleList
	return allRoleResponse, nil
}
