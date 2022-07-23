package internal

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	roleController "github.com/zvoleg/task-diary-back/internal/controllers/role"
	taskController "github.com/zvoleg/task-diary-back/internal/controllers/task"
	userController "github.com/zvoleg/task-diary-back/internal/controllers/user"
)

type Application struct {
	rout *mux.Router
	db   *sqlx.DB
}

func NewApplication(db *sqlx.DB) Application {
	return Application{
		rout: mux.NewRouter(),
		db:   db,
	}
}

func (app *Application) Init() {
	userController.RegisterHandlers(app.rout, app.db)
	roleController.RegisterHandlers(app.rout, app.db)
	taskController.RegisterHandlers(app.rout, app.db)
}

func (app *Application) Run() {
	http.ListenAndServe(":9090", app.rout)
}
