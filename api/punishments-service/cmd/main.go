package main

import (
	"github.com/sirupsen/logrus"

	"punishments-service/cmd/server"
	"punishments-service/init/config"
	"punishments-service/init/logger"
	"punishments-service/pkg/util/constants"
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
