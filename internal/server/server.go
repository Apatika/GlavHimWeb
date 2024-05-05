package server

import (
	"context"
	"glavhim-app/internal/config"
	"log"
	"net/http"
)

type Server struct {
	server *http.Server
}

func New(config config.Config, handler http.Handler) *Server {
	return &Server{
		server: &http.Server{
			Addr:         ":" + config.Server.Port,
			Handler:      handler,
			ReadTimeout:  config.Server.ReadTimeout,
			WriteTimeout: config.Server.WriteTimeout,
		},
	}
}

func (s *Server) Run() error {
	log.Printf("server run on %v", s.server.Addr)
	return s.server.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
