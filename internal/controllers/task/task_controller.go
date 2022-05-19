package task

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/zvoleg/task-diary-back/internal/models"
	"github.com/zvoleg/task-diary-back/internal/services"
)

type taskController struct {
	serv services.TaskService
}

func NewTaskController(serv services.TaskService) taskController {
	return taskController{serv: serv}
}

func (ctrl *taskController) Get() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		taskIdStr := ctx.Param("task_id")

		taskId, err := uuid.Parse(taskIdStr)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		user, err := ctrl.serv.Get(ctx.Request().Context(), taskId)
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		return ctx.JSON(http.StatusOK, user)
	}
}

func (ctrl *taskController) Create() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var task models.TaskRequest
		err := ctx.Bind(&task)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		taskResponse, err := ctrl.serv.Create(ctx.Request().Context(), task)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return ctx.JSON(http.StatusOK, taskResponse)
	}
}

func (ctrl *taskController) Update() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		taskIdStr := ctx.Param("task_id")

		taskId, err := uuid.Parse(taskIdStr)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		var task models.TaskRequest
		err = ctx.Bind(&task)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		taskResponse, err := ctrl.serv.Update(ctx.Request().Context(), taskId, task)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return ctx.JSON(http.StatusOK, taskResponse)
	}
}

func (ctrl *taskController) Delete() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		taskIdStr := ctx.Param("task_id")

		taskId, err := uuid.Parse(taskIdStr)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		err = ctrl.serv.Delete(ctx.Request().Context(), taskId)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return ctx.JSON(http.StatusNoContent, nil)
	}
}
