package user

import (
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	userRepo "github.com/zvoleg/task-diary-back/internal/repositories/user"
	userService "github.com/zvoleg/task-diary-back/internal/services/user"
)

func RegisterHandlers(rout *mux.Router, db *sqlx.DB) {
	userRepo := userRepo.NewUserRepository(db)
	userServ := userService.NewUserService(userRepo)
	userCtrl := NewUserController(userServ)

	rout.HandleFunc("/api/v1/user/{user_id}", userCtrl.Get).Methods("GET")
	rout.HandleFunc("/api/v1/user", userCtrl.Create).Methods("POST")
	rout.HandleFunc("/api/v1/user/{user_id}", userCtrl.Update).Methods("PUT")
	rout.HandleFunc("/api/v1/user/{user_id}", userCtrl.Delete).Methods("DELETE")
}
