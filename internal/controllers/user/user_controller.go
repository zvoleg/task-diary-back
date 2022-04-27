package user

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/zvoleg/task-diary-back/internal/models"
	"github.com/zvoleg/task-diary-back/internal/services"
)

type userController struct {
	service services.UserService
}

func NewUserController(service services.UserService) userController {
	return userController{service: service}
}

func (controller *userController) Get() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		userIdStr := ctx.Param("user_id")

		userId, err := uuid.Parse(userIdStr)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		user, err := controller.service.Get(ctx.Request().Context(), userId)
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		return ctx.JSON(http.StatusOK, user)
	}
}

func (controller *userController) Create() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		user := new(models.UserRequest)

		err := ctx.Bind(user)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		userResponse, err := controller.service.Create(ctx.Request().Context(), user)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return ctx.JSON(http.StatusOK, userResponse)
	}
}

func (controller *userController) Update() echo.HandlerFunc {
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
		userResponse, err := controller.service.Update(ctx.Request().Context(), userId, user)
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err)
		}
		return ctx.JSON(http.StatusOK, userResponse)
	}
}

func (controller *userController) Delete() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		userIdStr := ctx.Param("user_id")

		userId, err := uuid.Parse(userIdStr)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		err = controller.service.Delete(ctx.Request().Context(), userId)
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err)
		}
		return ctx.JSON(http.StatusNoContent, nil)
	}
}
