package server

import (
	"context"
	"github.com/nomadbala/crust/server/internal/config"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func New(cfg config.App, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:           ":" + cfg.Port,
			Handler:        handler,
			MaxHeaderBytes: cfg.MaxHeaderBytes,
			ReadTimeout:    cfg.ReadTimeout,
			WriteTimeout:   cfg.WriteTimeout,
		},
	}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
