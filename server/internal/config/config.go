package config

import (
	"os"
	"time"
)

var (
	DB_URL                 string
	REDIS_URL              string
	APP_PORT               string
	JWT_TOKEN_SIGNING_KEY  string
	JWT_TOKEN_EXPIRES_DATE time.Duration
)

const (
	ENV_DB_URL                     = "SUPABASE_DATABASE_URL"
	ENV_REDIS_URL                  = "TEST_REDIS_URL"
	ENV_APP_PORT                   = "APP_PORT"
	ENV_TOKEN_SIGNING_KEY          = "JWT_SIGNING_KEY"
	DEFAULT_TOKEN_EXPIRES_DURATION = 12 * time.Hour
)

func New() {
	DB_URL = os.Getenv(ENV_DB_URL)
	REDIS_URL = os.Getenv(ENV_REDIS_URL)
	JWT_TOKEN_SIGNING_KEY = os.Getenv(ENV_TOKEN_SIGNING_KEY)
	JWT_TOKEN_EXPIRES_DATE = DEFAULT_TOKEN_EXPIRES_DURATION
	APP_PORT = os.Getenv(ENV_APP_PORT)
}
