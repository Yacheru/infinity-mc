package producer

import (
	"github.com/sirupsen/logrus"
	"payments-service/init/logger"
	"payments-service/pkg/util/constants"

	"github.com/IBM/sarama"

	"payments-service/init/config"
)

type KafkaProducer struct {
	producer sarama.AsyncProducer
}

func NewKafkaProducer(cfg *config.Config) (*KafkaProducer, error) {
	kafkaConfig := sarama.NewConfig()

	kafkaConfig.Producer.Partitioner = sarama.NewRandomPartitioner
	kafkaConfig.Producer.RequiredAcks = sarama.WaitForAll
	producer, err := sarama.NewAsyncProducer([]string{cfg.KafkaBroker}, kafkaConfig)
	if err != nil {
		return nil, err
	}

	return &KafkaProducer{
		producer: producer,
	}, nil
}

func (p *KafkaProducer) PrepareMessage(message []byte, cfg *config.Config) error {
	logger.DebugF("prepare message: %s", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryKafka}, string(message))
	logger.DebugF("sent to topic: %v", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryKafka}, cfg.KafkaTopic)

	p.producer.Input() <- &sarama.ProducerMessage{
		Topic: cfg.KafkaTopic,
		Value: sarama.ByteEncoder(message),
	}

	return nil
}
