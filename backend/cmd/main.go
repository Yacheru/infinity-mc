package main

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"github.com/yacheru/infinity-mc.ru/backend"
	"github.com/yacheru/infinity-mc.ru/backend/configs"
	"github.com/yacheru/infinity-mc.ru/backend/pkg/handler"
	"github.com/yacheru/infinity-mc.ru/backend/pkg/repository"
	"github.com/yacheru/infinity-mc.ru/backend/pkg/service"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	if err := configs.InitConfig(); err != nil {
		log.Fatalf("Error reading config.json file, %s", err.Error())
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

	go func() {
		if err := srv.Run(viper.GetString("api.port"), handlers.InitRoutes()); err != nil {
			log.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	log.Printf("Listening on port %s", viper.GetString("api.port"))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Println("Shutting down server...")

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatalf("error occured while shutting down server, %s", err.Error())
	}

	if err := db.Close(); err != nil {
		log.Fatalf("error occured while closing database, %s", err.Error())
	}
}
