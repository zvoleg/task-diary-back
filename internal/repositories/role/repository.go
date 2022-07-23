package role

import (
	"context"
	"database/sql"

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
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, errors.Wrap(models.NewErrNotFoundInDb(identifier), "repo.Role.Get")
		default:
			return nil, errors.Wrap(err, "repo.Role.Get: struct scan error")
		}
	}
	role := roleDb.Map()
	return &role, nil
}

func (rep *roleRepository) GetList(ctx context.Context) (*models.AllRoleResponse, error) {
	rows, err := rep.db.QueryxContext(ctx, getListScript)
	if err != nil {
		return nil, errors.Wrap(err, "repo.Role.GetList")
	}

	roleList := make([]*models.RoleResponse, 0)

	for rows.Next() {
		roleDb := new(models.RoleDb)
		if err := rows.StructScan(roleDb); err != nil {
			return nil, errors.Wrap(err, "repo.Role.GetList: struct scan error")
		}
		role := roleDb.Map()
		roleList = append(roleList, &role)
	}

	allRoleResponse := models.AllRoleResponse{
		Roles: roleList,
	}
	return &allRoleResponse, nil
}
