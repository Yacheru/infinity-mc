package main

import (
	"github.com/spf13/viper"
	"github.com/yacheru/infinity-mc.ru/backend"
	"github.com/yacheru/infinity-mc.ru/backend/pkg/handler"
	"github.com/yacheru/infinity-mc.ru/backend/pkg/repository"
	"log"
)

func main() {
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err.Error())
	}

	_, err := repository.NewDatabaseDB(repository.Config{
		Host:    viper.GetString("db.psql.host"),
		Port:    viper.GetString("db.psql.port"),
		User:    viper.GetString("db.psql.user"),
		Pass:    viper.GetString("db.psql.pass"),
		SSLMode: viper.GetString("db.psql.sslmode"),
	})

	if err != nil {
		log.Fatalf("Error connecting to database, %s", err.Error())
	}

	srv := new(backend.Server)
	handlers := new(handler.Handler)

	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
