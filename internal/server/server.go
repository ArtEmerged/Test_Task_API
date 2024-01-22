package server

import (
	"net/http"
	"test_task/config"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(cfg *config.Config, handlers http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + cfg.Port,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		Handler:        handlers,
	}
	return s.httpServer.ListenAndServe()
}
