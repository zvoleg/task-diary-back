package user

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	userRepo "github.com/zvoleg/task-diary-back/internal/repositories/user"
	userService "github.com/zvoleg/task-diary-back/internal/services/user"
)

func RegisterHandlers(group *echo.Group, db *sqlx.DB) {
	userRepo := userRepo.NewUserRepository(db)
	userServ := userService.NewUserService(userRepo)
	userCtrl := NewUserController(userServ)

	usersGroup := group.Group("/users")

	usersGroup.GET("/:user_id", userCtrl.Get())
	usersGroup.POST("", userCtrl.Create())
	usersGroup.PUT("/:user_id", userCtrl.Update())
	usersGroup.DELETE("/:user_id", userCtrl.Delete())
}
