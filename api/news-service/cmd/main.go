package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"

	"news-service/init/config"
	"news-service/init/logger"
	"news-service/internal/server"
	"news-service/pkg/constants"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	cfg := &config.ServerConfig

	if err := config.InitConfig(); err != nil {
		fmt.Println(err.Error())
		cancel()
	}

	log := logger.InitLogger(cfg.APIDebug)

	app, err := server.NewServer(ctx, cfg, log)
	if err != nil {
		cancel()
	}
	logger.Info("server configured", constants.LoggerCMD)

	if app != nil {
		if err := app.Run(cfg); err != nil {
			logger.Error(err.Error(), constants.LoggerCMD)

			cancel()
		}
		logger.Info("server is running", constants.LoggerCMD)
	}

	<-ctx.Done()

	if app != nil {
		if err := app.Shutdown(ctx); err != nil {
			logger.Error(err.Error(), constants.LoggerCMD)
		}

		logger.Info("http-server shutdown", constants.LoggerCMD)
	}

	logger.Info("service shutdown", constants.LoggerCMD)
}
