package config

import (
	"os"
	"time"
)

type (
	Config struct {
		Database Database
		Redis    Redis
		Token    Token
		App      App
		Resend   Resend
	}

	Database struct {
		Url string
	}

	Redis struct {
		Url string
	}

	Token struct {
		SigningKey string
		Expires    time.Duration
	}

	App struct {
		Port           string
		MaxHeaderBytes int
		ReadTimeout    time.Duration
		WriteTimeout   time.Duration
	}

	Resend struct {
		ApiKey string
	}
)

const (
	ENV_DB_URL            = "DOCKER_POSTGRES_URL"
	ENV_REDIS_URL         = "DOCKER_REDIS_URL"
	ENV_APP_PORT          = "APP_PORT"
	ENV_TOKEN_SIGNING_KEY = "JWT_SIGNING_KEY"
)

const (
	DEFAULT_TOKEN_EXPIRES_DURATION  = 12 * time.Hour
	DEFAULT_SERVER_READ_TIMEOUT     = 10 * time.Second
	DEFAULT_SERVER_WRITE_TIMEOUTS   = 10 * time.Second
	DEFAULT_SERVER_MAX_HEADER_BYTES = 1 << 20
)

func New() (cfg Config, err error) {
	cfg.Database.Url = os.Getenv(ENV_DB_URL)
	cfg.Redis.Url = os.Getenv(ENV_REDIS_URL)

	cfg.App = App{
		Port:           os.Getenv(ENV_APP_PORT),
		MaxHeaderBytes: DEFAULT_SERVER_MAX_HEADER_BYTES,
		ReadTimeout:    DEFAULT_SERVER_READ_TIMEOUT,
		WriteTimeout:   DEFAULT_SERVER_WRITE_TIMEOUTS,
	}

	cfg.Token.SigningKey = os.Getenv(ENV_TOKEN_SIGNING_KEY)
	cfg.Token.Expires = DEFAULT_TOKEN_EXPIRES_DURATION

	cfg.Resend.ApiKey = os.Getenv("RESEND_API_KEY_2")

	return
}
