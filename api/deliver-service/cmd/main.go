package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"

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
	logger.Info("Configuration loaded", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryConfig})
}

func main() {
	logger.Info("Starting server...", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryConfig})

	r, err := setups.SetupRCONConnection()
	if err != nil {
		logger.Fatal("error setting up RCON connection", logrus.Fields{constants.LoggerCategory: constants.LoggerFile})
		os.Exit(1)
	}

	ctx := context.Background()
	err = consumer.NewConsumerGroup(ctx, r)
	if err != nil {
		logger.FatalF("consumer error: %s", logrus.Fields{constants.LoggerCategory: constants.LoggerFile}, err)
		os.Exit(1)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("Shutting down server...", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryConfig})

	for {

	}
}
