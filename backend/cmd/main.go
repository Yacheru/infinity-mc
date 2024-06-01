package main

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/yacheru/infinity-mc.ru/backend"
	"github.com/yacheru/infinity-mc.ru/backend/init/config"
	"github.com/yacheru/infinity-mc.ru/backend/init/databases"
	"github.com/yacheru/infinity-mc.ru/backend/internal/app/handler"
	"github.com/yacheru/infinity-mc.ru/backend/internal/app/repository"
	"github.com/yacheru/infinity-mc.ru/backend/internal/app/service"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.DateTime,
		PrettyPrint:     true,
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "@timestamp",
			logrus.FieldKeyLevel: "@level",
			logrus.FieldKeyMsg:   "@message",
			logrus.FieldKeyFunc:  "@caller",
			logrus.FieldKeyFile:  "@file",
		},
	})
	logrus.SetReportCaller(true)

	err := config.InitConfig()
	if err != nil {
		logrus.Fatalf("failed to init config: %s", err.Error())
	}

	mysql, psql := databases.InitDatabases()

	repo := repository.NewRepository(mysql, psql)
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)

	srv := new(backend.Server)

	go func() {
		if err := srv.Run(viper.GetString("api.port"), handlers.InitRoutes()); err != nil {
			logrus.Errorf("error occurred while running http server: %s", err.Error())
		}
	}()

	logrus.Infof("Listening on port %s", viper.GetString("api.port"))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Info("Shutting down server...")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occurred while shutting down server, %s", err.Error())
	}

	if err := mysql.Close(); err != nil {
		logrus.Errorf("error occurred while closing mysql database, %s", err.Error())
	}

	if err := psql.Close(); err != nil {
		logrus.Errorf("error occurred while closing psql database, %s", err.Error())
	}
}
