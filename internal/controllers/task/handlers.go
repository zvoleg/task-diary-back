package task

import (
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	taskRepo "github.com/zvoleg/task-diary-back/internal/repositories/task"
	taskServ "github.com/zvoleg/task-diary-back/internal/services/task"
)

func RegisterHandlers(rout *mux.Router, db *sqlx.DB) {
	taskRepo := taskRepo.NewTaskRepository(db)
	taskServ := taskServ.NewTaskService(taskRepo)
	taskCtrl := NewTaskController(taskServ)

	rout.HandleFunc("/api/vq/task/{task_id}", taskCtrl.Get).Methods("GET")
	rout.HandleFunc("/api/vq/task", taskCtrl.Create).Methods("POST")
	rout.HandleFunc("/api/vq/task/{task_id}", taskCtrl.Update).Methods("PUT")
	rout.HandleFunc("/api/vq/task/{task_id}", taskCtrl.Delete).Methods("DELETE")
}
