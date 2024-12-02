package config

import (
	constants2 "deliver-service/pkg/constants"
	"github.com/spf13/viper"

	"deliver-service/init/logger"
)

var ServerConfig Config

type Config struct {
	ServiceDebug bool `mapstructure:"SERVICE_DEBUG"`

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
		logger.ErrorF("failed loading config: %v", constants2.LoggerCategoryConfig, err)
		return constants2.ErrLoadConfig
	}

	err = viper.Unmarshal(&ServerConfig)
	if err != nil {
		logger.ErrorF("failed to parse env to config struct: %v", constants2.LoggerCategoryConfig, err)
		return constants2.ErrParseConfig
	}

	if ServerConfig.RCONPort == 0 || ServerConfig.RCONIp == "" || ServerConfig.RCONPass == "" ||
		len(ServerConfig.KafkaBrokers) == 0 || ServerConfig.KafkaConsumerGroup == "" || ServerConfig.KafkaTopic == "" {
		return constants2.ErrEmptyVar
	}

	return nil
}
