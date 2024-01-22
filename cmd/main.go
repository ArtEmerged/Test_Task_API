package main

import (
	"test_task/config"
	"test_task/internal/app"
)

func main() {
	cfg := config.InitConfig()
	app.RunServer(cfg)
}
