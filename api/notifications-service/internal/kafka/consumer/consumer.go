package consumer

import (
	"encoding/json"

	"github.com/IBM/sarama"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"

	"notifications-service/init/logger"
	"notifications-service/internal/discord"
	"notifications-service/internal/entities"
	"notifications-service/pkg/constants"
)

type Consumer struct {
	wh discord.Sender
}

func NewConsumer(wh *discord.Webhook) *Consumer {
	return &Consumer{wh: wh}
}

var validate = validator.New(validator.WithRequiredStructEnabled())

func (c *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	var message = new(entities.Message)

	for {
		select {
		case <-session.Context().Done():
			session.Commit()

			logger.Info("end of consumer work...", logrus.Fields{constants.LoggerCategory: constants.ConsumerCategory})

			return session.Context().Err()
		case msg, ok := <-claim.Messages():
			if !ok {
				logger.Debug("messages channel is closed", logrus.Fields{constants.Category: constants.ConsumerCategory})
			}

			if err := json.Unmarshal(msg.Value, message); err != nil {
				logger.ErrorF("error unmarshal message: %s", logrus.Fields{constants.Category: constants.ConsumerCategory}, err.Error())
				logger.DebugF("unmarshaled message: %s", logrus.Fields{constants.LoggerCategory: constants.ConsumerCategory}, string(msg.Value))

				session.MarkMessage(msg, "failed")
				continue
			}

			if err := validate.Struct(message); err != nil {
				logger.DebugF("invalid message structure: %s", logrus.Fields{constants.Category: constants.ConsumerCategory}, err.Error())

				session.MarkMessage(msg, "failed")
				continue
			}

			go func() {
				if err := c.wh.SendNotification(message); err != nil {
					logger.ErrorF("error send notification", logrus.Fields{constants.Category: constants.ConsumerCategory}, err.Error())
				}
			}()

			logger.DebugF("consumed message: %v", logrus.Fields{constants.LoggerCategory: constants.ConsumerCategory}, message)

			session.MarkMessage(msg, "success")
		}
	}
}

func (c *Consumer) Setup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (c *Consumer) Cleanup(session sarama.ConsumerGroupSession) error {
	return nil
}
