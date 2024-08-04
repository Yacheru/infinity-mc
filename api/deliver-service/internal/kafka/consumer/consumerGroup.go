package consumer

import (
	"context"

	"github.com/IBM/sarama"
	"github.com/gorcon/rcon"
	"github.com/sirupsen/logrus"

	cfg "deliver-service/init/config"
	"deliver-service/init/logger"
	r "deliver-service/internal/rcon"
	"deliver-service/pkg/util/constants"
)

func NewConsumerGroup(ctx context.Context, rcon *rcon.Conn) error {
	config := sarama.NewConfig()

	config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategyRange()}
	config.Consumer.Offsets.Initial = sarama.OffsetNewest

	consumerGroup, err := sarama.NewConsumerGroup(cfg.ServerConfig.KafkaBrokers, cfg.ServerConfig.KafkaConsumerGroup, config)
	if err != nil {
		return err
	}

	return Subscribe(ctx, cfg.ServerConfig.KafkaTopic, consumerGroup, rcon)
}

func Subscribe(ctx context.Context, topic string, consumerGroup sarama.ConsumerGroup, rcon *rcon.Conn) error {
	rc := r.NewRCON(rcon)
	consumer := NewConsumer(rc)

	go func() {
		logger.Info("consumer join the group...", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryKafka})
		if err := consumerGroup.Consume(ctx, []string{topic}, consumer); err != nil {
			logger.ErrorF("error consume: %s", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryKafka}, err.Error())
		}
		if ctx.Err() != nil {
			return
		}
	}()

	return nil
}
