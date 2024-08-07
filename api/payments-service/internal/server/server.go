package server

import (
	"errors"
	"fmt"
	"net/http"
	"payments-service/internal/server/http/middlewares"
	"payments-service/internal/server/http/routes"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"payments-service/init/config"
	"payments-service/init/logger"
	"payments-service/internal/kafka/producer"
	"payments-service/pkg/util/constants"
)

type Server struct {
	HttpServer *http.Server
	Producer   *producer.KafkaProducer
}

func NewServer() (*Server, error) {
	kafkaProducer, err := producer.NewKafkaProducer(&config.ServerConfig)
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
		if err := s.HttpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.ErrorF("Failed to listen and serve: %+v", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryServer}, err)
		}
	}()

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
