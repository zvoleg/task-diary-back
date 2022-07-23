package user

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/zvoleg/task-diary-back/internal/models"
	"github.com/zvoleg/task-diary-back/internal/services"
)

type userController struct {
	serv services.UserService
}

func NewUserController(service services.UserService) userController {
	return userController{serv: service}
}

func (ctrl *userController) Get() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		userIdStr := ctx.Param("user_id")

		userId, err := uuid.Parse(userIdStr)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		user, err := ctrl.serv.Get(ctx.Request().Context(), userId)
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		return ctx.JSON(http.StatusOK, user)
	}
}

func (ctrl *userController) Create() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		user := new(models.UserRequest)

		err := ctx.Bind(user)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		userResponse, err := ctrl.serv.Create(ctx.Request().Context(), user)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return ctx.JSON(http.StatusOK, userResponse)
	}
}

func (ctrl *userController) Update() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		userIdStr := ctx.Param("user_id")

		userId, err := uuid.Parse(userIdStr)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		user := new(models.UserRequest)
		err = ctx.Bind(user)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		userResponse, err := ctrl.serv.Update(ctx.Request().Context(), userId, user)
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err)
		}
		return ctx.JSON(http.StatusOK, userResponse)
	}
}

func (ctrl *userController) Delete() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		userIdStr := ctx.Param("user_id")

		userId, err := uuid.Parse(userIdStr)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		err = ctrl.serv.Delete(ctx.Request().Context(), userId)
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err)
		}
		return ctx.JSON(http.StatusNoContent, nil)
	}
}
