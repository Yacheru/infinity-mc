package main

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"

	"deliver-service/init/config"
	"deliver-service/init/logger"
	"deliver-service/internal/kafka/consumer"
	"deliver-service/pkg/util/constants"
	"deliver-service/pkg/util/setups"
)

func init() {
	if err := config.InitConfig(); err != nil {
		logger.Fatal(err.Error(), logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryConfig})
	}
	logger.Info("configuration loaded", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryConfig})
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	defer cancel()

	r, err := setups.SetupRCONConnection()
	if err != nil {
		logger.ErrorF("error setting up RCON connection: %s", logrus.Fields{constants.LoggerCategory: constants.LoggerFile}, err.Error())

		cancel()
	}

	if err := consumer.NewConsumerGroup(ctx, r); err != nil {
		logger.ErrorF("consumer error: %s", logrus.Fields{constants.LoggerCategory: constants.LoggerFile}, err.Error())

		cancel()
	}

	<-ctx.Done()

	logger.Info("shutting down server...", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryConfig})
}
