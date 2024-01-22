package routes

import (
	"github.com/ssss-tantalum/echo-kick/internal/core"
	"github.com/ssss-tantalum/echo-kick/internal/handler/api"
	"github.com/ssss-tantalum/echo-kick/internal/handler/view"
)

func InitRoutes(app *core.App) {
	echo := app.Router()

	// View Routes
	indexHandler := view.NewIndexHandler()
	echo.GET("/", indexHandler.IndexView)

	// API Rputes
	a := echo.Group("/api")
	userHandler := api.NewUserHadler(app)
	a.GET("/user", userHandler.ShowUser)
}
