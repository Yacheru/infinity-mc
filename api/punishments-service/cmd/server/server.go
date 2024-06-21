package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"punishments-service/init/config"
	"punishments-service/init/logger"
	"punishments-service/internal/http/middlewares"
	"punishments-service/internal/http/routes"
	"punishments-service/pkg/util/constants"
	"punishments-service/pkg/util/setups"
)

type Server struct {
	HttpServer *http.Server
}

func NewServer() (*Server, error) {
	db, err := setups.SetupMysqlDatabase()
	if err != nil {
		return nil, err
	}

	router := setupRouter()

	api := router.Group("/punishments")
	routes.NewPunishmentsRoute(api, db).Routes()

	server := &http.Server{
		Addr:              fmt.Sprintf(":%d", config.ServerConfig.APIPort),
		Handler:           router,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		MaxHeaderBytes:    1 << 20,
	}

	return &Server{
		HttpServer: server,
	}, nil
}

func (s *Server) Run() error {
	go func() {
		logger.InfoF("success to listen and serve on :%d", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryServer}, config.ServerConfig.APIPort)
		if err := s.HttpServer.ListenAndServe(); err != nil && !errors.Is(http.ErrServerClosed, err) {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	logger.Info("shutdown server ...", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryServer})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.HttpServer.Shutdown(ctx); err != nil {
		return fmt.Errorf("error when shutdown server: %v", err)
	}

	<-ctx.Done()
	logger.Info("timeout of 5 seconds.", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryServer})
	logger.Info("server exiting", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryServer})

	return nil
}

func setupRouter() *gin.Engine {
	var mode = gin.ReleaseMode
	if config.ServerConfig.APIDebug {
		mode = gin.DebugMode
	}
	gin.SetMode(mode)

	router := gin.New()

	router.Use(middlewares.CORSMiddleware())
	router.Use(gin.LoggerWithFormatter(logger.HTTPLogger))
	router.Use(gin.Recovery())

	return router
}
