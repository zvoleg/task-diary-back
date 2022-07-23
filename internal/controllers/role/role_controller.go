package role

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/zvoleg/task-diary-back/internal/services"
)

type roleController struct {
	serv services.RoleService
}

func NewRoleController(serv services.RoleService) roleController {
	return roleController{serv: serv}
}

func (ctrl *roleController) Get(w http.ResponseWriter, r *http.Request) {
	roleIdStr := mux.Vars(r)["role_id"]

	roleId, err := uuid.Parse(roleIdStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	role, err := ctrl.serv.Get(roleId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	roleJsonStr, _ := json.Marshal(role)
	w.WriteHeader(http.StatusOK)
	w.Write(roleJsonStr)
}

func (ctrl *roleController) GetList(w http.ResponseWriter, r *http.Request) {
	allRoles, err := ctrl.serv.GetList()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	allRolesJsonStr, _ := json.Marshal(allRoles)
	w.WriteHeader(http.StatusOK)
	w.Write(allRolesJsonStr)
}
