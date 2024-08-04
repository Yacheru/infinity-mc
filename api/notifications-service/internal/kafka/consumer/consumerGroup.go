package consumer

import (
	"context"
	"github.com/IBM/sarama"
	"github.com/sirupsen/logrus"
	"notifications-service/init/config"
	"notifications-service/init/logger"
	"notifications-service/internal/discord"
	"notifications-service/pkg/constants"
)

func NewConsumerGroup(ctx context.Context, cfg *config.Config) error {
	kafkaConfig := sarama.NewConfig()

	kafkaConfig.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategyRange()}
	kafkaConfig.Consumer.Offsets.Initial = sarama.OffsetNewest

	consumerGroup, err := sarama.NewConsumerGroup(cfg.KafkaBrokers, cfg.KafkaConsumerGroup, kafkaConfig)
	if err != nil {
		return err
	}

	return Subscribe(ctx, cfg, consumerGroup)
}

func Subscribe(ctx context.Context, cfg *config.Config, consumerGroup sarama.ConsumerGroup) error {
	webhook := discord.NewWebhookClient(cfg)
	consumer := NewConsumer(webhook)

	go func() {
		logger.Info("consumer join the group...", logrus.Fields{constants.Category: constants.ConsumerCategory})
		if err := consumerGroup.Consume(ctx, []string{cfg.KafkaTopic}, consumer); err != nil {
			logger.ErrorF("error consume: %v", logrus.Fields{constants.Category: constants.ConsumerCategory}, err.Error())
		}
		if ctx.Err() != nil {
			return
		}
	}()

	return nil
}
