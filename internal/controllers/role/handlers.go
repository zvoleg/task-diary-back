package role

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	roleRepo "github.com/zvoleg/task-diary-back/internal/repositories/role"
	roleServ "github.com/zvoleg/task-diary-back/internal/services/role"
)

func RegisterHandlers(group *echo.Group, db *sqlx.DB) {
	roleRepo := roleRepo.NewRoleRepository(db)
	roleServ := roleServ.NewRoleService(roleRepo)
	roleCtrl := NewRoleController(roleServ)

	roleGroup := group.Group("/roles")

	roleGroup.GET("/:role_id", roleCtrl.Get())
	roleGroup.GET("", roleCtrl.GetList())
}
