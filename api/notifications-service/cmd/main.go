package main

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"

	"notifications-service/init/config"
	"notifications-service/init/logger"
	"notifications-service/internal/kafka/consumer"
	"notifications-service/pkg/constants"
)

func init() {
	if err := config.InitConfig(); err != nil {
		logger.Fatal(constants.ErrFailedLoadConfig.Error(), logrus.Fields{constants.Category: constants.CmdLogger})
	}
	logger.Info("successfully loaded config!", logrus.Fields{constants.Category: constants.CmdLogger})
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	defer cancel()

	if err := consumer.NewConsumerGroup(ctx, &config.ServerConfig); err != nil {
		logger.Error(err.Error(), logrus.Fields{constants.Category: constants.CmdLogger})

		cancel()
	}

	<-ctx.Done()

	logger.Info("service exited.", logrus.Fields{constants.Category: constants.CmdLogger})
}
