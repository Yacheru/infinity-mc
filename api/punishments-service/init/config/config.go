package config

import (
	"github.com/spf13/viper"
	"punishments-service/pkg/constants"

	"punishments-service/init/logger"
)

var ServerConfig Config

type Config struct {
	APIPort        int    `mapstructure:"API_PORT"`
	APIEnvironment string `mapstructure:"API_ENVIRONMENT"`
	APIDebug       bool   `mapstructure:"API_DEBUG"`
	ApiEntry       string `mapstructure:"API_ENTRY"`

	MySQLDb   string `mapstructure:"MYSQL_DATABASE"`
	MySQLUser string `mapstructure:"MYSQL_USER"`
	MySQLPass string `mapstructure:"MYSQL_PASSWORD"`
	MySQLAddr string `mapstructure:"MYSQL_ADDR"`
	MySQLPort string `mapstructure:"MYSQL_PORT"`
	MySQLUrl  string `mapstructure:"MYSQL_URL"`

	JWTSalt string `mapstructure:"JWT_SALT"`
}

func InitConfig() error {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath("./configs")
	viper.AddConfigPath("/")
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

	if ServerConfig.APIPort == 0 || ServerConfig.APIEnvironment == "" || ServerConfig.MySQLUrl == "" {
		logger.Error("empty requirement variable!", constants.LoggerCategoryConfig)
		return constants.ErrEmptyVar
	}

	return nil
}
