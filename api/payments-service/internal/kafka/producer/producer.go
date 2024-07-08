package producer

import (
	"github.com/sirupsen/logrus"
	"payments-service/init/logger"
	"payments-service/pkg/util/constants"

	"github.com/IBM/sarama"

	cfg "payments-service/init/config"
)

type KafkaProducer struct {
	producer sarama.AsyncProducer
}

func NewKafkaProducer() (*KafkaProducer, error) {
	config := sarama.NewConfig()

	config.Version = sarama.V2_6_0_0
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	producer, err := sarama.NewAsyncProducer(cfg.ServerConfig.KafkaBrokers, config)
	if err != nil {
		return nil, err
	}

	return &KafkaProducer{
		producer: producer,
	}, nil
}

func (p *KafkaProducer) PrepareMessage(message []byte) error {
	logger.DebugF("message: %s", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryKafka}, message)
	logger.DebugF("send to topic: %v", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryKafka}, cfg.ServerConfig.KafkaTopic)

	p.producer.Input() <- &sarama.ProducerMessage{
		Topic: cfg.ServerConfig.KafkaTopic,
		Value: sarama.ByteEncoder(message),
	}

	return nil
}
