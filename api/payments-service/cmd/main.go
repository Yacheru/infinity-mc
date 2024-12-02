package main

import (
	"context"
	"os/signal"
	"payments-service/pkg/constants"
	"syscall"

	"payments-service/init/config"
	"payments-service/init/logger"
	"payments-service/internal/server"
)

func init() {
	if err := config.InitConfig(); err != nil {
		logger.Error(err.Error(), constants.LoggerCategoryConfig)
	}
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	defer cancel()

	cfg := &config.ServerConfig

	log := logger.InitLogger(cfg.APIDebug)

	app, err := server.NewServer(ctx, cfg, log)
	if err != nil {
		cancel()
	}

	if app != nil {
		if err := app.Run(); err != nil {
			cancel()
		}
	}

	logger.Info("service started", constants.LoggerCategoryServer)

	<-ctx.Done()

	logger.Info("service exiting...", constants.LoggerCategoryServer)

	if app != nil {
		if err := app.Shutdown(ctx); err != nil {
			logger.Error(err.Error(), constants.LoggerCategoryServer)
		}
	}

	logger.Info("server exited", constants.LoggerCategoryServer)
}
