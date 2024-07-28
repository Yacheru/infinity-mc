package main

import (
	"github.com/sirupsen/logrus"
	"payments-service/internal/server"

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
	app, err := server.NewServer()
	if err != nil {
		logger.Fatal(err.Error(), logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryServer})
	}
	if err := app.Run(); err != nil {
		logger.Fatal(err.Error(), logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryServer})
	}
}
