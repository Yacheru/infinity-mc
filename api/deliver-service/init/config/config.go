package config

import (
	"deliver-service/init/logger"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"deliver-service/pkg/util/constants"
)

var ServerConfig Config

type Config struct {
	RCONPort int    `mapstructure:"RCON_PORT"`
	RCONIp   string `mapstructure:"RCON_IP"`
	RCONPass string `mapstructure:"RCON_PASS"`

	KafkaBrokers       []string `mapstructure:"KAFKA_BROKERS"`
	KafkaConsumerGroup string   `mapstructure:"KAFKA_CONSUMER_GROUP"`
	KafkaTopic         string   `mapstructure:"KAFKA_TOPIC"`
}

func InitConfig() error {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath("./configs")
	viper.AddConfigPath("/")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		logger.FatalF("failed loading config: %v", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryConfig}, err)

		return constants.ErrLoadConfig
	}

	err = viper.Unmarshal(&ServerConfig)
	if err != nil {
		logger.FatalF("failed to parse env to config struct: %v", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryConfig}, err)

		return constants.ErrParseConfig
	}

	if ServerConfig.RCONPort == 0 || ServerConfig.RCONIp == "" || ServerConfig.RCONPass == "" ||
		len(ServerConfig.KafkaBrokers) == 0 || ServerConfig.KafkaConsumerGroup == "" || ServerConfig.KafkaTopic == "" {
		return constants.ErrEmptyVar
	}

	return nil
}
