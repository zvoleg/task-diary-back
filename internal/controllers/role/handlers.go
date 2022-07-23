package role

import (
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	roleRepo "github.com/zvoleg/task-diary-back/internal/repositories/role"
	roleServ "github.com/zvoleg/task-diary-back/internal/services/role"
)

func RegisterHandlers(rout *mux.Router, db *sqlx.DB) {
	roleRepo := roleRepo.NewRoleRepository(db)
	roleServ := roleServ.NewRoleService(roleRepo)
	roleCtrl := NewRoleController(roleServ)

	rout.HandleFunc("/api/v1/role/{role_id}", roleCtrl.Get).Methods("GET")
	rout.HandleFunc("/api/v1/role", roleCtrl.GetList).Methods("GET")
}
