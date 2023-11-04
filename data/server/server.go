package server

import (
	"apatikaWebServer/data/config"
	"context"
	"log"
	"net/http"
)

type Server struct {
	server *http.Server
}

func New() *Server {
	return &Server{
		server: &http.Server{
			Addr:         ":" + config.Config.Server.Port,
			Handler:      nil,
			ReadTimeout:  config.Config.Server.ReadTimeout,
			WriteTimeout: config.Config.Server.WriteTimeout,
		},
	}
}

func (s *Server) Run() error {
	log.Println("start server")
	return s.server.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
