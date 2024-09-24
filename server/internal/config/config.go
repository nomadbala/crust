package config

import (
	"errors"
	"github.com/joho/godotenv"
	"github.com/nomadbala/crust/server/pkg/log"
	"go.uber.org/zap"
	"os"
	"time"
)

type (
	Configs struct {
		DB    DB
		Redis Redis
		App   App
		Token TokenConfig
	}

	DB struct {
		Url string
	}

	Redis struct {
		Url string
	}

	App struct {
		Port string
	}

	TokenConfig struct {
		SigningKey string
		Expires    time.Duration
	}
)

const (
	ENV_DB_URL                     = "SUPABASE_DATABASE_URL"
	ENV_REDIS_URL                  = "TEST_REDIS_URL"
	ENV_APP_PORT                   = "APP_PORT"
	ENV_TOKEN_SIGNING_KEY          = "JWT_SIGNING_KEY"
	DEFAULT_APP_PORT               = "8080"
	DEFAULT_TOKEN_EXPIRES_DURATION = 12 * time.Hour
)

func New() (cfg Configs, err error) {
	if err := godotenv.Load(); err != nil {
		log.Logger.Error("error occurred while loading .env file", zap.Error(err))
	}

	cfg.DB.Url = os.Getenv(ENV_DB_URL)
	cfg.Redis.Url = os.Getenv(ENV_REDIS_URL)
	cfg.Token.SigningKey = os.Getenv(ENV_TOKEN_SIGNING_KEY)
	cfg.Token.Expires = DEFAULT_TOKEN_EXPIRES_DURATION
	cfg.App.Port = os.Getenv(ENV_APP_PORT)

	if cfg.App.Port == "" {
		cfg.App.Port = DEFAULT_APP_PORT
	}

	if cfg.Token.SigningKey == "" {
		return Configs{}, errors.New("error occurred while loading .env token signin key")
	}

	if cfg.Redis.Url == "" {
		return Configs{}, errors.New("error occurred while loading .env redis url")
	}

	if cfg.DB.Url == "" {
		return Configs{}, errors.New("error occurred while loading .env redis url")
	}

	return cfg, nil
}
