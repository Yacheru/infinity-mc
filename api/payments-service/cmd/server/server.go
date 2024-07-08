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

	"payments-service/init/config"
	"payments-service/init/logger"
	"payments-service/internal/http/middlewares"
	"payments-service/internal/http/routes"
	"payments-service/internal/kafka/producer"
	"payments-service/pkg/util/constants"
)

type Server struct {
	HttpServer *http.Server
	Producer   *producer.KafkaProducer
}

func NewServer() (*Server, error) {
	kafkaProducer, err := producer.NewKafkaProducer()
	if err != nil {
		return nil, err
	}

	router := setupRouter()

	api := router.Group("/payments")
	routes.NewPaymentsRoute(api, kafkaProducer).Routes()

	server := &http.Server{
		Addr:              fmt.Sprintf(":%d", config.ServerConfig.APIPort),
		Handler:           router,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		MaxHeaderBytes:    1 << 20, // 1mb
	}

	return &Server{
		HttpServer: server,
		Producer:   kafkaProducer,
	}, nil
}

func (s *Server) Run() error {
	go func() {
		logger.InfoF("success to listen and serve on :%d port", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryServer}, config.ServerConfig.APIPort)
		if err := s.HttpServer.ListenAndServe(); err != nil && !errors.Is(http.ErrServerClosed, err) {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	logger.Info("shutdown server in 5 seconds...", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryServer})

	if err := s.HttpServer.Shutdown(ctx); err != nil {
		logrus.Errorf("error when shutdown http server: %v", err)

		return err
	}

	//if err := s.Producer.Close(); err != nil {
	//	logrus.Errorf("error when close kafka producer: %v", err)
	//
	//	return err
	//}

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
