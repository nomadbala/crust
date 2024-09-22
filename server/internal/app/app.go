package app

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/nomadbala/crust/server/db/postgres/sqlc"
	"github.com/nomadbala/crust/server/internal/handler"
	"github.com/nomadbala/crust/server/internal/repository"
	"github.com/nomadbala/crust/server/internal/service"
	"github.com/nomadbala/crust/server/pkg/log"
	"github.com/nomadbala/crust/server/pkg/resend"
	"github.com/nomadbala/crust/server/pkg/server"
	"go.uber.org/zap"
	"os"
)

var ctx = context.Background()

func Run() {
	log.ConfigureLogger()
	defer func(Logger *zap.Logger) {
		err := Logger.Sync()
		if err != nil {
			log.Logger.Error("failed to sync logger", zap.Error(err))
		}
	}(log.Logger)

	if err := godotenv.Load(); err != nil {
		log.Logger.Error("error occurred while loading .env file", zap.Error(err))
	}

	conn, err := pgx.Connect(ctx, os.Getenv("POSTGRES_DATABASE_URL"))
	if err != nil {
		log.Logger.Error("failed to connect to database", zap.Error(err))
	}

	defer func(conn *pgx.Conn, ctx context.Context) {
		err := conn.Close(ctx)
		if err != nil {
			log.Logger.Error("failed to close connection", zap.Error(err))
		}
	}(conn, ctx)

	repos := repository.New(sqlc.New(conn), ctx)
	services := service.New(repos)
	handlers := handler.New(services)

	resend.ConfigureResendClient()

	servers := server.New(os.Getenv("APP_PORT"), handlers.ConfigureRoutes())

	if err := servers.Run(); err != nil {
		log.Logger.Error("failed to start server", zap.Error(err))
	}
}
