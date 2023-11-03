package server

import (
	"apatikaWebServer/internal/config"
	"apatikaWebServer/pkg/router"
	"context"
	"log"
	"net/http"
)

type Server struct {
	server *http.Server
}

func New() *Server {
	handler := router.New()
	return &Server{
		server: &http.Server{
			Addr:         ":" + config.Config.Server.Port,
			Handler:      handler,
			ReadTimeout:  config.Config.Server.ReadTimeout,
			WriteTimeout: config.Config.Server.WriteTimeout,
		},
	}
}

func (s *Server) Run() error {
	log.Println("server started")
	return s.server.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
