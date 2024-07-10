package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	
	"punishments-service/init/logger"
	"punishments-service/pkg/util/constants"
)

var ServerConfig Config

type Config struct {
	APIPort        int    `mapstructure:"API_PORT"`
	APIEnvironment string `mapstructure:"API_ENVIRONMENT"`
	APIDebug       bool   `mapstructure:"API_DEBUG"`

	MYSQLDriver string `mapstructure:"MYSQL_DRIVER"`
	MYSQLURL    string `mapstructure:"MYSQL_URL"`
	MYSQLDsn    string `mapstructure:"MYSQL_DSN"`
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
		ServerConfig.MYSQLDriver == "" || ServerConfig.MYSQLURL == "" || ServerConfig.MYSQLDsn == "" {
		logger.Error("empty requirement variable!", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryConfig})

		return constants.ErrEmptyVar
	}

	return nil
}
