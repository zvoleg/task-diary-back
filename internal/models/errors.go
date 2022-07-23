package models

import (
	"fmt"

	"github.com/google/uuid"
)

type ErrNotFoundInDb struct {
	identifier uuid.UUID
}

func NewErrNotFoundInDb(identifier uuid.UUID) ErrNotFoundInDb {
	return ErrNotFoundInDb{identifier: identifier}
}

func (e ErrNotFoundInDb) Error() string {
	return fmt.Sprintf("id '%s' not founded", e.identifier)
}
