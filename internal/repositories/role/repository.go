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
	roleDb := new(models.RoleDb)

	if err := rep.db.QueryRowxContext(
		ctx,
		getScript,
		identifier,
	).StructScan(roleDb); err != nil {
		return nil, errors.Wrap(err, "roleRepo.Get: struct scan error")
	}
	role := models.RoleResponse{
		RoleId:      roleDb.RoleId,
		Name:        roleDb.Name,
		Description: roleDb.Description,
	}
	return &role, nil
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
		roleDb := new(models.RoleDb)
		if err := rows.StructScan(roleDb); err != nil {
			return nil, errors.Wrap(err, "roleRepo.GetList: struct scan error")
		}
		role := models.RoleResponse{
			RoleId:      roleDb.RoleId,
			Name:        roleDb.Name,
			Description: roleDb.Description,
		}
		roleList = append(roleList, &role)
	}

	allRoleResponse := new(models.AllRoleResponse)
	allRoleResponse.Roles = roleList
	return allRoleResponse, nil
}
