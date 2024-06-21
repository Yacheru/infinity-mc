package config

import (
	"github.com/spf13/viper"

	"deliver-service/pkg/util/constants"
)

var ServerConfig Config

type Config struct {
	APIPort        int    `mapstructure:"API_PORT"`
	APIEnvironment string `mapstructure:"API_ENVIRONMENT"`
	APIDebug       bool   `mapstructure:"API_DEBUG"`

	RCONPort int    `mapstructure:"RCON_PORT"`
	RCONIp   string `mapstructure:"RCON_IP"`
	RCONPass string `mapstructure:"RCON_PASS"`
}

func InitConfig() error {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath("./configs")
	viper.AddConfigPath("/")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return constants.ErrLoadConfig
	}

	err = viper.Unmarshal(&ServerConfig)
	if err != nil {
		return constants.ErrParseConfig
	}

	if ServerConfig.APIPort == 0 || ServerConfig.APIEnvironment == "" ||
		ServerConfig.RCONPort == 0 || ServerConfig.RCONIp == "" || ServerConfig.RCONPass == "" {
		return constants.ErrEmptyVar
	}

	return nil
}
