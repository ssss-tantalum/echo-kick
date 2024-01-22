package main

import (
	"fmt"

	"github.com/ssss-tantalum/echo-kick/internal/core"
	"github.com/ssss-tantalum/echo-kick/internal/routes"
)

func main() {
	cfg := core.NewAppConfig()
	app := core.NewApp(cfg)

	routes.InitRoutes(app)
	echo := app.Router()

	port := fmt.Sprintf(":%d", cfg.HttpPort)
	echo.Logger.Fatal(echo.Start(port))
}
