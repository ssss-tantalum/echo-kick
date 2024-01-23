package core

import (
	"database/sql"
	"log"
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
	"github.com/uptrace/bun/extra/bundebug"
)

type App struct {
	router *echo.Echo
	config *AppConfig

	dbOnce sync.Once
	db     *bun.DB
}

func NewApp(config *AppConfig) *App {
	app := &App{
		config: config,
	}
	app.initRouter()

	return app
}

func (app *App) IsDebug() bool {
	return app.config.Debug
}

func (app *App) Config() *AppConfig {
	return app.config
}

func (app *App) Router() *echo.Echo {
	return app.router
}

func (app *App) DB() *bun.DB {
	app.dbOnce.Do(func() {
		sqldb, err := sql.Open(sqliteshim.ShimName, app.config.DbDatabase)
		if err != nil {
			log.Fatal(err)
		}

		db := bun.NewDB(sqldb, sqlitedialect.New())
		db.AddQueryHook(bundebug.NewQueryHook(
			bundebug.WithEnabled(app.IsDebug()),
			bundebug.FromEnv(""),
		))

		app.db = db
	})

	return app.db
}
