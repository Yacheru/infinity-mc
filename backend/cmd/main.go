package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"github.com/yacheru/infinity-mc.ru/backend"
	"github.com/yacheru/infinity-mc.ru/backend/configs"
	"github.com/yacheru/infinity-mc.ru/backend/pkg/handler"
	"github.com/yacheru/infinity-mc.ru/backend/pkg/repository"
	"github.com/yacheru/infinity-mc.ru/backend/pkg/service"
	"log"
)

// main запускаем api
func main() {
	if err := configs.InitConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err.Error())
	}

	db, err := repository.NewDatabaseDB(repository.Config{
		Host:    viper.GetString("db.mdb.host"),
		Port:    viper.GetString("db.mdb.port"),
		User:    viper.GetString("db.mdb.user"),
		Pass:    viper.GetString("db.mdb.pass"),
		Dbname:  viper.GetString("db.mdb.dbname"),
		SSLMode: viper.GetString("db.mdb.sslmode"),
	})

	if err != nil {
		log.Fatalf("Error connecting to database, %s", err.Error())
	}

	repo := repository.NewRepository(db)
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)

	srv := new(backend.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}

}
