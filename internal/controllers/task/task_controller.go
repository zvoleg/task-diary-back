package task

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/zvoleg/task-diary-back/internal/models"
	"github.com/zvoleg/task-diary-back/internal/services"
)

type taskController struct {
	serv services.TaskService
}

func NewTaskController(serv services.TaskService) taskController {
	return taskController{serv: serv}
}

func (ctrl *taskController) Get(w http.ResponseWriter, r *http.Request) {
	taskIdStr := mux.Vars(r)["task_id"]

	taskId, err := uuid.Parse(taskIdStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	task, err := ctrl.serv.Get(taskId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	taskJsonStr, _ := json.Marshal(task)
	w.WriteHeader(http.StatusOK)
	w.Write(taskJsonStr)
}

func (ctrl *taskController) Create(w http.ResponseWriter, r *http.Request) {
	var task models.TaskRequest
	jsonDecoder := json.NewDecoder(r.Body)
	err := jsonDecoder.Decode(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	taskResponse, err := ctrl.serv.Create(task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	taskJsonStr, _ := json.Marshal(taskResponse)
	w.WriteHeader(http.StatusCreated)
	w.Write(taskJsonStr)
}

func (ctrl *taskController) Update(w http.ResponseWriter, r *http.Request) {
	taskIdStr := mux.Vars(r)["task_id"]

	taskId, err := uuid.Parse(taskIdStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var task models.TaskRequest
	jsonDecoder := json.NewDecoder(r.Body)
	err = jsonDecoder.Decode(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	taskResponse, err := ctrl.serv.Update(taskId, task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	taskJsonStr, _ := json.Marshal(taskResponse)
	w.WriteHeader(http.StatusOK)
	w.Write(taskJsonStr)
}

func (ctrl *taskController) Delete(w http.ResponseWriter, r *http.Request) {
	taskIdStr := mux.Vars(r)["task_id"]

	taskId, err := uuid.Parse(taskIdStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = ctrl.serv.Delete(taskId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
