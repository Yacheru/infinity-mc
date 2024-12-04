package producer

import (
	"github.com/IBM/sarama"
	"payments-service/init/logger"
	"payments-service/pkg/constants"

	"payments-service/init/config"
)

//go:generate go run github.com/vektra/mockery/v2@v2.49.1 --name=Producer
type Producer interface {
	PrepareMessage(message []byte) error
}

type KafkaProducer struct {
	producer sarama.AsyncProducer
	topic    string
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
		topic:    cfg.KafkaTopic,
	}, nil
}

func (p *KafkaProducer) PrepareMessage(message []byte) error {
	logger.DebugF("prepare message: %s", constants.LoggerCategoryKafka, string(message))

	p.producer.Input() <- &sarama.ProducerMessage{
		Topic: p.topic,
		Value: sarama.ByteEncoder(message),
	}

	logger.DebugF("sent to topic: %v", constants.LoggerCategoryKafka, p.topic)

	return nil
}
