package server

import (
	"apatikaWebServer/internal/config"
	"context"
	"net/http"
)

type Server struct {
	srv *http.Server
}

func New(cfg *config.Config, handler http.Handler) *Server {
	return &Server{
		srv: &http.Server{
			Addr:         ":" + cfg.Server.Port,
			Handler:      handler,
			ReadTimeout:  cfg.Server.ReadTimeout,
			WriteTimeout: cfg.Server.WriteTimeout,
		},
	}
}

func (s *Server) Run() error {
	return s.srv.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}
