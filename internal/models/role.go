package models

import (
	"github.com/google/uuid"
)

type RoleDb struct {
	RoleId      uuid.UUID `db:"role_id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
}

type RoleResponse struct {
	RoleId      uuid.UUID `json:"role_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

type AllRoleResponse struct {
	Roles []*RoleResponse `json:"data"`
}
