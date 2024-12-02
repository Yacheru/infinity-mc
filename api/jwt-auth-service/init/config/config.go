package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"jwt-auth-service/pkg/constants"
)

var ServerConfig Config

type Config struct {
	ApiPort  int    `mapstructure:"API_PORT"`
	ApiDebug bool   `mapstructure:"API_DEBUG"`
	ApiEntry string `mapstructure:"API_ENTRY"`

	PostgresDSN string `mapstructure:"POSTGRESQL_DSN"`

	Salt string `mapstructure:"SALT"`

	AccessTokenTTL  int `mapstructure:"ACCESS_TOKEN_TTL"`
	RefreshTokenTTL int `mapstructure:"REFRESH_TOKEN_TTL"`

	EmailSender         string `mapstructure:"EMAIL_SENDER"`
	EmailSenderPassword string `mapstructure:"EMAIL_SENDER_PASSWORD"`

	SmtpHost string `mapstructure:"SMTP_SERVER"`
	SmtpPort string `mapstructure:"SMTP_SERVER_PORT"`

	MongoDB         string `mapstructure:"MONGODB_DATABASE"`
	MongoCollection string `mapstructure:"MONGODB_COLLECTION"`
	MongoURI        string `mapstructure:"MONGODB_URI"`
	MongoTTL        int32  `mapstructure:"MONGODB_TTL"`

	RedisHost     string `mapstructure:"REDIS_HOST"`
	RedisPassword string `mapstructure:"REDIS_PASSWORD"`
	RedisDatabase int    `mapstructure:"REDIS_DATABASE"`
	RedisTTL      int    `mapstructure:"REDIS_TTL"`
}

func InitConfig() error {
	viper.SetConfigType("env")
	viper.SetConfigName(".env")
	viper.AddConfigPath("./configs")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		logrus.Error(err.Error(), logrus.WithFields(logrus.Fields{constants.LoggerCategory: constants.ConfigCategory}))
		return err
	}

	if err := viper.Unmarshal(&ServerConfig); err != nil {
		logrus.Error(err.Error(), logrus.WithFields(logrus.Fields{constants.LoggerCategory: constants.ConfigCategory}))
		return err
	}

	if err := checkVars(); err != nil {
		logrus.Error(err.Error(), logrus.WithFields(logrus.Fields{constants.LoggerCategory: constants.ConfigCategory}))
		return err
	}

	return nil
}

func checkVars() error {
	if ServerConfig.ApiPort == 0 || ServerConfig.ApiEntry == "" {
		return constants.ApiVarsRequiredError
	}

	if ServerConfig.PostgresDSN == "" {
		return constants.PostgresDSNRequiredError
	}

	return nil
}
