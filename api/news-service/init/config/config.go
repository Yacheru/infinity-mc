package config

import (
	"errors"
	"github.com/disgoorg/snowflake/v2"
	"github.com/spf13/viper"

	"news-service/init/logger"
	"news-service/pkg/constants"
)

var ServerConfig Config

type Config struct {
	APIPort  int    `mapstructure:"API_PORT"`
	APIDebug bool   `mapstructure:"API_DEBUG"`
	APIEntry string `mapstructure:"API_ENTRY"`

	PostgresDSN string `mapstructure:"POSTGRESQL_DSN"`

	ElasticClient   string `mapstructure:"ELASTICSEARCH_CLIENT"`
	ElasticPassword string `mapstructure:"ELASTICSEARCH_PASSWORD"`
	ElasticUsername string `mapstructure:"ELASTICSEARCH_USERNAME"`
	ElasticIndex    string `mapstructure:"ELASTICSEARCH_INDEX"`

	RedisHost     string `mapstructure:"REDIS_HOST"`
	RedisPassword string `mapstructure:"REDIS_PASSWORD"`
	RedisTTL      int    `mapstructure:"REDIS_TTL"`

	WebhookID    snowflake.ID `mapstructure:"WEBHOOK_ID"`
	WebhookToken string       `mapstructure:"WEBHOOK_TOKEN"`
}

func InitConfig() error {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath("./configs")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(&ServerConfig); err != nil {
		return err
	}

	if err := checkVars(); err != nil {
		return err
	}

	return nil
}

func checkVars() error {
	if ServerConfig.APIPort == 0 || ServerConfig.APIEntry == "" {
		logger.Error(errors.New(constants.EmptyRequiredVar.Error()+": api").Error(), constants.LoggerConfig)

		return constants.EmptyRequiredVar
	}

	if ServerConfig.PostgresDSN == "" {
		logger.Error(errors.New(constants.EmptyRequiredVar.Error()+": postgres").Error(), constants.LoggerConfig)

		return constants.EmptyRequiredVar
	}

	if ServerConfig.ElasticClient == "" || ServerConfig.ElasticPassword == "" || ServerConfig.ElasticUsername == "" || ServerConfig.ElasticIndex == "" {
		logger.Error(errors.New(constants.EmptyRequiredVar.Error()+": elastic").Error(), constants.LoggerConfig)

		return constants.EmptyRequiredVar
	}

	if ServerConfig.RedisHost == "" || ServerConfig.RedisPassword == "" {
		logger.Error(errors.New(constants.EmptyRequiredVar.Error()+": redis").Error(), constants.LoggerConfig)

		return constants.EmptyRequiredVar
	}

	if ServerConfig.WebhookID == 0 || ServerConfig.WebhookToken == "" {
		logger.Error(errors.New(constants.EmptyRequiredVar.Error()+": webhook").Error(), constants.LoggerConfig)

		return constants.EmptyRequiredVar
	}

	return nil
}
