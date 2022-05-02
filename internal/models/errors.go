package models

import "fmt"

var EntityNotFoundInDb = fmt.Errorf("sql: no rows in result set")

type ErrNotFoundInDb struct {
	Msg string
}

func (e *ErrNotFoundInDb) Error() string {
	return e.Msg
}
