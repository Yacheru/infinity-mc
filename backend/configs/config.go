package configs

import "github.com/spf13/viper"

// InitConfig Инициализирует работу с конфигурационным файлом
func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
