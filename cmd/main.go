package main

import (
	"log"
	"test_task/config"
	"test_task/internal/app"
)

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	app.RunServer(cfg)
}

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)
