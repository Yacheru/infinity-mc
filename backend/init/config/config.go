package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

// InitConfig Инициализирует работу с конфигурационным файлом
func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigType("yml")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		logrus.Printf("Error reading config file, %s", err.Error())
		os.Exit(1)
	}

	return viper.ReadInConfig()
}
