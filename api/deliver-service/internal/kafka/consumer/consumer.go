package consumer

import (
	"deliver-service/internal/entities"
	"encoding/json"

	"github.com/IBM/sarama"
	"github.com/sirupsen/logrus"

	"deliver-service/init/logger"
	"deliver-service/internal/rcon"
	"deliver-service/pkg/util/constants"
)

type Consumer struct {
	rcon rcon.Deliver
}

func NewConsumer(rcon rcon.Deliver) *Consumer {
	return &Consumer{
		rcon: rcon,
	}
}

func (c *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	var mc = new(entities.MC)

	for {
		select {
		case <-session.Context().Done():
			session.Commit()

			logger.Info("end of consumer work...", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryKafka})

			return nil
		case msg, ok := <-claim.Messages():
			if !ok {
				logger.Debug("message channel is closed", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryKafka})
			}

			err := json.Unmarshal(msg.Value, &mc)
			if err != nil {
				logger.ErrorF("error unmarshal message: %v", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryKafka}, err)
				logger.DebugF("unmarshaled message: %s", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryKafka}, string(msg.Value))

				return err
			}

			err = c.rcon.DeliverService(mc.Nickname, mc.Service, mc.Duration)
			if err != nil {
				logger.ErrorF("error deliver service: %v", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryKafka}, err)
				return err
			}

			logger.DebugF("consumed message: %v", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryKafka}, mc)

			session.MarkMessage(msg, "")
		}
	}
}

func (c *Consumer) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (c *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}
