package config

import (
	"github.com/disgoorg/snowflake/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"notifications-service/pkg/constants"

	"notifications-service/init/logger"
)

var ServerConfig Config

type Config struct {
	KafkaBrokers       []string `mapstructure:"KAFKA_BROKERS"`
	KafkaConsumerGroup string   `mapstructure:"KAFKA_CONSUMER_GROUP"`
	KafkaTopic         string   `mapstructure:"KAFKA_TOPIC"`

	WebhookToken string       `mapstructure:"WEBHOOK_TOKEN"`
	WebhookID    snowflake.ID `mapstructure:"WEBHOOK_ID"`
}

func InitConfig() error {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath("./configs")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		logger.Error(err.Error(), logrus.Fields{constants.Category: constants.ConfigCategory})

		return constants.ErrLoadConfig
	}

	if err := viper.Unmarshal(&ServerConfig); err != nil {
		logger.Error(err.Error(), logrus.Fields{constants.Category: constants.ConfigCategory})

		return constants.ErrParseConfig
	}

	if len(ServerConfig.KafkaBrokers) == 0 || ServerConfig.KafkaConsumerGroup == "" || ServerConfig.KafkaTopic == "" {
		logger.Error(constants.ErrEmptyVar.Error(), logrus.Fields{constants.Category: constants.ConfigCategory})

		return constants.ErrEmptyVar
	}

	return nil
}
