package core

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (app *App) initRouter() {
	app.router = echo.New()

	app.router.Static("/web/static", "web/static")
	app.router.File("/favicon.ico", "favicon.ico")
	app.router.Use(middleware.Recover())
	app.router.Use(middleware.Secure())
	app.router.Use(middleware.CORS())
	app.router.Use(middleware.Logger())
}

// TODO: write some middlewares, like error handler
