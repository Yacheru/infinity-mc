package main

import (
	"context"
	"deliver-service/internal/utils"
	"deliver-service/pkg/constants"
	"os/signal"
	"syscall"

	"deliver-service/init/config"
	"deliver-service/init/logger"
	"deliver-service/internal/kafka/consumer"
)

func init() {
	if err := config.InitConfig(); err != nil {
		logger.Fatal(err.Error(), constants.LoggerCategoryConfig)
	}
	logger.Info("configuration loaded", constants.LoggerCategoryConfig)
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	defer cancel()

	cfg := &config.ServerConfig
	logger.InitLogger(cfg.ServiceDebug)

	r, err := utils.SetupRCONConnection()
	if err != nil {
		logger.ErrorF("error setting up RCON connection: %s", constants.LoggerFile, err.Error())

		cancel()
	}

	if err := consumer.NewConsumerGroup(ctx, r); err != nil {
		logger.ErrorF("consumer error: %s", constants.LoggerFile, err.Error())

		cancel()
	}

	<-ctx.Done()

	logger.Info("shutting down server...", constants.LoggerCategoryConfig)
}
