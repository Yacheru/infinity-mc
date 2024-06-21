package main

import (
	"github.com/sirupsen/logrus"

	"deliver-service/init/config"
	"deliver-service/init/logger"
	"deliver-service/pkg/util/constants"
)

func init() {
	if err := config.InitConfig(); err != nil {
		logger.Fatal(err.Error(), logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryConfig})
	}
	logger.Info("configuration loaded", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryConfig})
}

func main() {

}
