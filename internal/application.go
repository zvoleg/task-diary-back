package internal

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	roleController "github.com/zvoleg/task-diary-back/internal/controllers/role"
	taskController "github.com/zvoleg/task-diary-back/internal/controllers/task"
	userController "github.com/zvoleg/task-diary-back/internal/controllers/user"
)

type Application struct {
	echo *echo.Echo
	db   *sqlx.DB
}

func NewApplication(db *sqlx.DB) Application {
	return Application{
		echo: echo.New(),
		db:   db,
	}
}

func (app *Application) Init() {
	v1Group := app.echo.Group("/api/v1")

	userController.RegisterHandlers(v1Group, app.db)
	roleController.RegisterHandlers(v1Group, app.db)
	taskController.RegisterHandlers(v1Group, app.db)
}

func (app *Application) Run() {
	app.echo.Start(":9090")
}
