package models

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type TaskType int

const (
	task TaskType = iota + 1
	root
	bug
	note
)

var toString = map[TaskType]string{
	task: "task",
	root: "root",
	bug:  "bug",
	note: "note",
}

var toId = map[string]TaskType{
	"task": task,
	"root": root,
	"bug":  bug,
	"note": note,
}

// Implementation of sql.Scaner interface
func (t *TaskType) Scan(value interface{}) error {
	s := value.(string)
	*t = toId[s]
	return nil
}

// Implementation of sql.driver.Valuer interface
func (t TaskType) Value() (driver.Value, error) {
	return toString[t], nil
}

func (t TaskType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(fmt.Sprintf("\"%s\"", toString[t]))
	return buffer.Bytes(), nil
}

func (t *TaskType) UnmarshalJSON(b []byte) error {
	var str string
	err := json.Unmarshal(b, &str)
	if err != nil {
		return errors.Wrap(err, "models.TaskType.UnmarshalJSON: can't read json value")
	}
	*t = toId[str]
	return nil
}

type TaskRequest struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Type        TaskType `json:"type"`
	Status      string   `json:"status"`
	Tags        string   `json:"tags"`
}

type TaskDb struct {
	TaskId      uuid.UUID  `db:"task_id"`
	BoardId     uuid.UUID  `db:"board_id"`
	Title       string     `db:"title"`
	Description string     `db:"description"`
	Type        TaskType   `db:"type"`
	Status      string     `db:"status"`
	AuthorId    uuid.UUID  `db:"author_id"`
	CreatedAt   time.Time  `db:"created_at"`
	UpdatedAt   *time.Time `db:"updated_at"`
	Tags        string     `db:"tags"`
	IsDeleted   bool       `db:"is_deleted"`
}

type TaskResponse struct {
	TaskId      uuid.UUID  `json:"task_id"`
	BoardId     uuid.UUID  `json:"board_id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Type        TaskType   `json:"type"`
	Status      string     `json:"status"`
	AuthorId    uuid.UUID  `json:"author_id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	Tags        string     `json:"tags"`
}

type AllTaskResponse struct {
	Tasks []*TaskResponse `json:"data"`
}

func (taskDb *TaskDb) Map() TaskResponse {
	return TaskResponse{
		TaskId:      taskDb.TaskId,
		BoardId:     taskDb.BoardId,
		Title:       taskDb.Title,
		Description: taskDb.Description,
		Type:        taskDb.Type,
		Status:      taskDb.Status,
		AuthorId:    taskDb.AuthorId,
		CreatedAt:   taskDb.CreatedAt,
		UpdatedAt:   taskDb.UpdatedAt,
		Tags:        taskDb.Tags,
	}
}
