package user

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/zvoleg/task-diary-back/internal/models"
	"github.com/zvoleg/task-diary-back/internal/services"
)

type userController struct {
	serv services.UserService
}

func NewUserController(service services.UserService) userController {
	return userController{serv: service}
}

func (ctrl *userController) Get(w http.ResponseWriter, r *http.Request) {
	userIdStr := mux.Vars(r)["user_id"]

	userId, err := uuid.Parse(userIdStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := ctrl.serv.Get(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	jsonUserStr, _ := json.Marshal(user)
	w.WriteHeader(http.StatusOK)
	w.Write(jsonUserStr)
}

func (ctrl *userController) Create(w http.ResponseWriter, r *http.Request) {
	user := new(models.UserRequest)

	jsonDecoder := json.NewDecoder(r.Body)
	err := jsonDecoder.Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	userResponse, err := ctrl.serv.Create(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	jsonUserStr, _ := json.Marshal(userResponse)
	w.WriteHeader(http.StatusOK)
	w.Write(jsonUserStr)
}

func (ctrl *userController) Update(w http.ResponseWriter, r *http.Request) {
	userIdStr := mux.Vars(r)["user_id"]

	userId, err := uuid.Parse(userIdStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user := new(models.UserRequest)
	jsonDecoder := json.NewDecoder(r.Body)
	err = jsonDecoder.Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	userResponse, err := ctrl.serv.Update(userId, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	jsonUserStr, _ := json.Marshal(userResponse)
	w.WriteHeader(http.StatusOK)
	w.Write(jsonUserStr)
}

func (ctrl *userController) Delete(w http.ResponseWriter, r *http.Request) {
	userIdStr := mux.Vars(r)["user_id"]

	userId, err := uuid.Parse(userIdStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = ctrl.serv.Delete(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
