package role

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
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
	role := models.RoleResponse{}

	rep.db.QueryRowxContext(
		ctx,
		getScript,
		identifier,
	).StructScan(role)
	panic("unimplemented")
}

func (rep *roleRepository) GetList(context.Context) (*models.AllRoleResponse, error) {
	panic("unimplemented")
}
