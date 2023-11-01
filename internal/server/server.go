package server

import (
	"apatikaWebServer/internal/config"
	"context"
	"net/http"
)

type Server struct {
	srv *http.Server
}

func New(cfg *config.HTTPconfig, handler http.Handler) *Server {
	return &Server{
		srv: &http.Server{
			Addr:           ":" + cfg.Port,
			Handler:        handler,
			ReadTimeout:    cfg.ReadTimeout,
			WriteTimeout:   cfg.WriteTimeout,
			MaxHeaderBytes: cfg.MaxHeaderMegabytes << 20,
		},
	}
}

func (s *Server) Run() error {
	return s.srv.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}
