package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"os/signal"
	"payments-service/internal/server"
	"syscall"

	"payments-service/init/config"
	"payments-service/init/logger"
	"payments-service/pkg/util/constants"
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

	app, err := server.NewServer()
	if err != nil {
		logger.Error(err.Error(), logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryServer})

		cancel()
	}
	if err := app.Run(); err != nil {
		logger.Error(err.Error(), logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryServer})

		cancel()
	}

	<-ctx.Done()

	logger.Info("server exiting", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryServer})
}
