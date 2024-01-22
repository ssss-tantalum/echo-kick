package main

import (
	"fmt"

	"github.com/labstack/echo/v4/middleware"
	"github.com/ssss-tantalum/echo-kick/internal/core"
	"github.com/ssss-tantalum/echo-kick/internal/handler"
)

func main() {
	app := core.New()

	echo := app.Echo
	echo.Static("/web/static", "web/static")
	echo.Use(middleware.Logger())

	initRoutes(app)

	echo.Logger.Fatal(echo.Start(fmt.Sprintf(":%d", app.HttpPort)))
}

func initRoutes(app *core.App) {
	echo := app.Echo

	userHandler := handler.NewUserHandler(app)
	echo.GET("/user", userHandler.HandleUserShow)
}
