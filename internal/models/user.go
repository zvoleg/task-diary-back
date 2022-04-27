package models

import (
	"time"

	"github.com/google/uuid"
)

type UserRequest struct {
	Name        string     `json:"name" db:"name"`
	Description string     `json:"description" db:"description"`
	RoleId      *uuid.UUID `json:"role_id" db:"role_id"`
}

type UserDb struct {
	UserId      uuid.UUID  `db:"user_id"`
	AccountId   *uuid.UUID `db:"account_id"`
	Name        string     `db:"name"`
	Description string     `db:"description"`
	RoleId      *uuid.UUID `db:"role_id"`
	CreatedAt   time.Time  `db:"created_at"`
	UpdatedAt   *time.Time `db:"updated_at"`
	IsDeleted   bool       `db:"is_deleted"`
}

type UserResponse struct {
	UserId      uuid.UUID  `json:"user_id"`
	AccountId   *uuid.UUID `json:"account_id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	RoleId      *uuid.UUID `json:"role_id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

type AllUserResponse struct {
	Users []*UserResponse `json:"data"`
}
