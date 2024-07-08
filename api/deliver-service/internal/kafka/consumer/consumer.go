package consumer

import (
	"encoding/json"

	"github.com/IBM/sarama"
	"github.com/sirupsen/logrus"

	"deliver-service/init/logger"
	"deliver-service/internal/rcon"
	"deliver-service/pkg/util/constants"
)

type Consumer struct {
	r rcon.Deliver
}

func NewConsumer(rcon rcon.Deliver) *Consumer {
	return &Consumer{
		r: rcon,
	}
}

func (c *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	var mc rcon.MC

	for message := range claim.Messages() {
		select {
		case <-session.Context().Done():
			session.Commit()
			return nil
		default:
			err := json.Unmarshal(message.Value, &mc)
			if err != nil {
				logger.ErrorF("error unmarshal message: %v", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryKafka}, err)
				logger.DebugF("message: %v", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryKafka}, message.Value)

				return err
			}

			err = c.r.DeliverService(mc.Nickname, mc.Service, mc.Duration)
			if err != nil {
				logger.ErrorF("error deliver service: %v", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryKafka}, err)

				return err
			}
			session.MarkMessage(message, "")

			logger.InfoF("Consumed message: %v", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryKafka}, mc)
		}
	}

	return nil
}

func (c *Consumer) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (c *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}
