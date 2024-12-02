package main

import (
	"context"
	"os"
	"os/signal"
	"punishments-service/internal/server"
	"punishments-service/pkg/constants"
	"syscall"

	"punishments-service/init/config"
	"punishments-service/init/logger"
)

func init() {
	if err := config.InitConfig(); err != nil {
		logger.Error(err.Error(), constants.LoggerCategoryConfig)
	}
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	cfg := &config.ServerConfig

	_ = logger.InitLogger(cfg.APIDebug)

	app, err := server.NewServer(ctx, cfg)
	if err != nil {
		logger.Error(err.Error(), constants.LoggerCategoryServer)
		cancel()
	}

	if app != nil {
		if err := app.Run(ctx); err != nil {
			cancel()
		}
	}

	logger.InfoF("success to listen and serve on :%d (pid: %d)", constants.LoggerCategoryServer, config.ServerConfig.APIPort, os.Getpid())

	<-ctx.Done()

	logger.Info("service exiting...", constants.LoggerCategoryServer)

	if app != nil {
		if err := app.Shutdown(ctx); err != nil {
			logger.Error(err.Error(), constants.LoggerCategoryServer)
		}
	}

	logger.Info("service exited", constants.LoggerCategoryServer)
}
