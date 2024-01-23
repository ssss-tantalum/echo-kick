package core

import (
	"log"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

type AppConfig struct {
	Debug bool `env:"DEGUG"`

	HttpPort      int    `env:"HTTP_PORT"`
	SessionSecret string `env:"SESSION_SECRET"`

	DbDatabase string `env:"DB_DATABASE"`
}

func NewAppConfig() *AppConfig {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	var cfg AppConfig
	if err := env.Parse(&cfg); err != nil {
		panic(err)
	}

	return &AppConfig{
		Debug:         cfg.Debug,
		HttpPort:      cfg.HttpPort,
		SessionSecret: cfg.SessionSecret,
		DbDatabase:    cfg.DbDatabase,
	}
}
