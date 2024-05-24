package main

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"github.com/yacheru/infinity-mc.ru/backend"
	"github.com/yacheru/infinity-mc.ru/backend/init/config"
	"github.com/yacheru/infinity-mc.ru/backend/init/logger"
	"github.com/yacheru/infinity-mc.ru/backend/pkg/handler"
	"github.com/yacheru/infinity-mc.ru/backend/pkg/repository"
	"github.com/yacheru/infinity-mc.ru/backend/pkg/service"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	if err := config.InitConfig(); err != nil {
		fmt.Printf("Error reading config.json file, %s", err.Error())
	}

	log := logger.SetupLogger(viper.GetString("status"))

	db, err := repository.NewDatabaseDB(repository.Config{
		Host:    viper.GetString("db.mdb.host"),
		Port:    viper.GetString("db.mdb.port"),
		User:    viper.GetString("db.mdb.user"),
		Pass:    viper.GetString("db.mdb.pass"),
		Dbname:  viper.GetString("db.mdb.dbname"),
		SSLMode: viper.GetString("db.mdb.sslmode"),
	})

	if err != nil {
		log.Error("Error connecting to database", err.Error())
	}

	repo := repository.NewRepository(db)
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)

	srv := new(backend.Server)

	go func() {
		if err := srv.Run(viper.GetString("api.port"), handlers.InitRoutes(log)); err != nil {
			log.Error("error occured while running http server: %s", err.Error())
		}
	}()

	log.Info("Listening on port " + viper.GetString("api.port"))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Info("Shutting down server...")

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Error("error occurred while shutting down server, " + err.Error())
	}

	if err := db.Close(); err != nil {
		log.Error("error occurred while closing database, " + err.Error())
	}
}
