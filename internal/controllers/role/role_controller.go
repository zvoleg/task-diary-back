package role

import (
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/zvoleg/task-diary-back/internal/services"
)

type roleController struct {
	serv services.RoleService
}

func NewRoleController(serv services.RoleService) roleController {
	return roleController{serv: serv}
}

func (ctrl *roleController) Get() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		roleIdStr := ctx.Param("role_id")

		roleId, err := uuid.Parse(roleIdStr)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		role, err := ctrl.serv.Get(ctx.Request().Context(), roleId)
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound, errors.Unwrap(err).Error())
		}
		return ctx.JSON(http.StatusOK, role)
	}
}

func (ctrl *roleController) GetList() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		allRoles, err := ctrl.serv.GetList(ctx.Request().Context())
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound, errors.Unwrap(err).Error())
		}
		return ctx.JSON(http.StatusOK, allRoles)
	}
}
