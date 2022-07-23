package task

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	taskRepo "github.com/zvoleg/task-diary-back/internal/repositories/task"
	taskServ "github.com/zvoleg/task-diary-back/internal/services/task"
)

func RegisterHandlers(group *echo.Group, db *sqlx.DB) {
	taskRepo := taskRepo.NewTaskRepository(db)
	taskServ := taskServ.NewTaskService(taskRepo)
	taskCtrl := NewTaskController(taskServ)

	taskGroup := group.Group("/task")

	taskGroup.GET("/:task_id", taskCtrl.Get())
	taskGroup.POST("", taskCtrl.Create())
	taskGroup.PUT("/:task_id", taskCtrl.Update())
	taskGroup.DELETE("/:task_id", taskCtrl.Delete())
}
