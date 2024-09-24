package app

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/nomadbala/crust/server/db/postgres/sqlc"
	"github.com/nomadbala/crust/server/internal/cache"
	"github.com/nomadbala/crust/server/internal/config"
	"github.com/nomadbala/crust/server/internal/handler"
	"github.com/nomadbala/crust/server/internal/repository"
	"github.com/nomadbala/crust/server/internal/service"
	"github.com/nomadbala/crust/server/pkg/log"
	"github.com/nomadbala/crust/server/pkg/resend"
	"github.com/nomadbala/crust/server/pkg/server"
	"github.com/nomadbala/crust/server/pkg/store"
	"go.uber.org/zap"
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

	configs, err := config.New()
	if err != nil {
		return
	}

	conn, err := pgx.Connect(ctx, configs.DB.Url)
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

	redisClient := store.RedisClient{}
	redisClient.New(ctx, configs.Redis.Url)

	caches := cache.New(&redisClient)

	services := service.New(repos, caches)
	handlers := handler.New(services)

	resend.ConfigureResendClient()

	servers := server.New(configs.App.Port, handlers.ConfigureRoutes())

	if err := servers.Run(); err != nil {
		log.Logger.Error("failed to start server", zap.Error(err))
	}
}
