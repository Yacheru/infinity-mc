package config

import (
	"errors"
	"github.com/spf13/viper"
	"websocket-chat-service/init/logger"
	"websocket-chat-service/pkg/constants"
)

var ServerConfig Config

type Config struct {
	ApiDebug bool   `mapstructure:"API_DEBUG"`
	ApiPort  int    `mapstructure:"API_PORT"`
	ApiEntry string `mapstructure:"API_ENTRY"`

	RedisAddr     string `mapstructure:"REDIS_ADDR"`
	RedisPassword string `mapstructure:"REDIS_PASSWORD"`

	ScyllaHosts    []string `mapstructure:"SCYLLA_HOSTS"`
	ScyllaKeyspace string   `mapstructure:"SCYLLA_KEYSPACE"`

	WebsocketURL   string `mapstructure:"WEBSOCKET_URL"`
	WebsocketLimit int    `mapstructure:"WEBSOCKET_LIMIT"`

	BearerAuth string `mapstructure:"BEARER_AUTH"`
}

func InitConfig() error {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath("./configs")
	if err := viper.ReadInConfig(); err != nil {
		logger.Error(err.Error(), constants.ConfigCategory)
		return err
	}

	if err := viper.Unmarshal(&ServerConfig); err != nil {
		logger.Error(err.Error(), constants.ConfigCategory)
		return err
	}

	if err := CheckVars(); err != nil {
		logger.Error(err.Error(), constants.ConfigCategory)
		return err
	}

	return nil
}

func CheckVars() error {
	if ServerConfig.ApiPort == 0 || ServerConfig.ApiEntry == "" {
		return errors.New("API_PORT and API_ENTRY is required environment variables")
	}

	if ServerConfig.RedisAddr == "" || ServerConfig.RedisPassword == "" {
		return errors.New("REDIS_ADDR and REDIS_PASSWORD are required environment variables")
	}

	if ServerConfig.ScyllaKeyspace == "" || len(ServerConfig.ScyllaHosts) == 0 {
		return errors.New("SCYLLA_KEYSPACE and SCYLLA_HOSTS is required environment variable")
	}

	if ServerConfig.WebsocketURL == "" || ServerConfig.BearerAuth == "" {
		return errors.New("WEBSOCKET_URL and BEARER_AUTH are required environment variables")
	}

	return nil
}
