package core

import (
	"database/sql"
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
	"github.com/uptrace/bun/extra/bundebug"
)

type App struct {
	Echo     *echo.Echo
	Env      string
	HttpPort int

	dbOnce sync.Once
	db     *bun.DB
}

func New() *App {
	if err := godotenv.Load(); err != nil {
		log.Println("Failed to load .env file, continuing with default values")
	}

	httpPort, _ := strconv.Atoi(GetDefaultEnv("HTTP_PORT", "3000"))
	env := GetDefaultEnv("ENV", "development")

	app := &App{
		Echo:     echo.New(),
		Env:      env,
		HttpPort: httpPort,
	}

	return app
}

func (app *App) DB() *bun.DB {
	dbName := GetDefaultEnv("DB_DATABASE", "app.db")

	app.dbOnce.Do(func() {
		sqldb, err := sql.Open(sqliteshim.ShimName, dbName)
		if err != nil {
			log.Panicf("Failed to connect database, Error: %s", err)
		}

		db := bun.NewDB(sqldb, sqlitedialect.New())
		db.AddQueryHook(bundebug.NewQueryHook(
			bundebug.WithEnabled(false),
			bundebug.FromEnv(""),
		))

		app.db = db
	})

	return app.db
}

func GetDefaultEnv(key, defaultValue string) string {
	if value, exsits := os.LookupEnv(key); exsits {
		return value
	}
	return defaultValue
}
