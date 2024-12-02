package config

import (
	"github.com/spf13/viper"
	"payments-service/pkg/constants"

	"payments-service/init/logger"
)

var ServerConfig Config

type Config struct {
	APIPort        int    `mapstructure:"API_PORT"`
	APIEnvironment string `mapstructure:"API_ENVIRONMENT"`
	APIDebug       bool   `mapstructure:"API_DEBUG"`

	YKassaID   string `mapstructure:"YKASSA_ID"`
	YKassaPass string `mapstructure:"YKASSA_PASS"`

	KafkaBroker string `mapstructure:"KAFKA_BROKER"`
	KafkaTopic  string `mapstructure:"KAFKA_TOPIC"`

	PostgresDSN string `mapstructure:"POSTGRESQL_DSN"`
}

func InitConfig() error {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath("./configs")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		logger.ErrorF("error read in config: %v", constants.LoggerCategoryConfig, err)

		return constants.ErrLoadConfig
	}

	err = viper.Unmarshal(&ServerConfig)
	if err != nil {
		logger.ErrorF("error unmarshal config: %v", constants.LoggerCategoryConfig, err)

		return constants.ErrParseConfig
	}

	if ServerConfig.APIPort == 0 || ServerConfig.APIEnvironment == "" ||
		ServerConfig.YKassaID == "" || ServerConfig.YKassaPass == "" ||
		ServerConfig.KafkaBroker == "" || ServerConfig.KafkaTopic == "" {
		logger.Error("empty requirement variable!", constants.LoggerCategoryConfig)

		return constants.ErrEmptyVar
	}

	return err
}
