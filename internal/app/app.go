package app

import (
	"log"
	"test_task/config"
	"test_task/internal/handler"
	"test_task/internal/repository"
	"test_task/internal/server"
	"test_task/internal/service"
)

func RunServer(cfg *config.Config) {

	db, err := repository.InitPostgresDB(cfg.DB)
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(server.Server)

	err = srv.Run(cfg, handlers.InitRouter())

}
