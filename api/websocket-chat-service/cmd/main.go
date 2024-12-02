package main

import (
	"context"
	"os/signal"
	"syscall"

	"websocket-chat-service/init/config"
	"websocket-chat-service/init/logger"
	"websocket-chat-service/internal/server"
	"websocket-chat-service/pkg/constants"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL)

	cfg := &config.ServerConfig

	if err := config.InitConfig(); err != nil {
		logger.Error(err.Error(), constants.MainCategory)
		cancel()
	}

	logger.InitLogger(cfg.ApiDebug)

	app, err := server.NewServer(ctx, cfg)
	if err != nil {
		cancel()
	}

	if app != nil {
		app.Run()
	}
	logger.Info("service is running", constants.MainCategory)

	<-ctx.Done()

	if app != nil {
		if err := app.Shutdown(ctx); err != nil {
			logger.Error(err.Error(), constants.MainCategory)
		}
	}

	logger.Info("service is shutdown", constants.MainCategory)
}
