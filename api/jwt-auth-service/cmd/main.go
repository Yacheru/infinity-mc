package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"jwt-auth-service/init/config"
	"jwt-auth-service/init/logger"
	"jwt-auth-service/internal/server"
	"jwt-auth-service/pkg/constants"
	"net/http"
	"os/signal"
	"syscall"
)

// @title Jwt-Auth-Api
// @version 1.0
// @description jwt-auth-service

// @host localhost
// @BasePath /

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	cfg := &config.ServerConfig

	if err := config.InitConfig(); err != nil {
		fmt.Println(err.Error())
		cancel()
	}

	log := logger.InitLogger(cfg.ApiDebug)

	app, err := server.NewServer(ctx, cfg, log)
	if err != nil {
		cancel()
	}
	logger.Info("server configured", constants.MainCategory)

	if app != nil {
		errs, gCtx := errgroup.WithContext(ctx)
		errs.Go(func() error {
			return app.Run()
		})

		errs.Go(func() error {
			<-gCtx.Done()
			return app.Shutdown(gCtx)
		})

		if err := errs.Wait(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error(err.Error(), constants.MainCategory)
			cancel()
		}
	}

	<-ctx.Done()

	logger.Info("service shutdown", constants.MainCategory)
}
