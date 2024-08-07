package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"payments-service/init/logger"
	"payments-service/pkg/util/constants"
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

	PSQLDriver string `mapstructure:"POSTGRES_DRIVER"`
	PSQLURL    string `mapstructure:"POSTGRES_URL"`
	PSQLDsn    string `mapstructure:"POSTGRES_DSN"`
}

func InitConfig() error {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath("./configs")
	viper.AddConfigPath("/")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		logger.ErrorF("error read in config: %v", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryConfig}, err)

		return constants.ErrLoadConfig
	}

	err = viper.Unmarshal(&ServerConfig)
	if err != nil {
		logger.ErrorF("error unmarshal config: %v", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryConfig}, err)

		return constants.ErrParseConfig
	}

	if ServerConfig.APIPort == 0 || ServerConfig.APIEnvironment == "" ||
		ServerConfig.YKassaID == "" || ServerConfig.YKassaPass == "" ||
		ServerConfig.PSQLDriver == "" || ServerConfig.PSQLURL == "" || ServerConfig.PSQLDsn == "" {
		logger.Error("empty requirement variable!", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryConfig})

		return constants.ErrEmptyVar
	}

	return err
}
