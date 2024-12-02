package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"payments-service/init/config"
	"payments-service/init/logger"
	"payments-service/internal/repository/postgres"
	"payments-service/internal/server/http/routes"
	"payments-service/pkg/constants"
)

type Server struct {
	HttpServer *http.Server
}

func NewServer(ctx context.Context, cfg *config.Config, log *logrus.Logger) (*Server, error) {
	db, err := postgres.InitPostgresqlConnection(ctx, cfg, log)
	if err != nil {
		return nil, err
	}

	router := setupRouter()
	api := router.Group("/pay")
	routes.NewPaymentsRoute(api, db).Routes(cfg.APIDebug)

	server := &http.Server{
		Addr:              fmt.Sprintf(":%d", config.ServerConfig.APIPort),
		Handler:           router,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		MaxHeaderBytes:    1 << 20, // 1mb
	}

	return &Server{
		HttpServer: server,
	}, nil
}

func (s *Server) Run() error {
	go func() {
		logger.InfoF("success to listen and serve on :%d port", constants.LoggerCategoryServer, config.ServerConfig.APIPort)
		if err := s.HttpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.ErrorF("Failed to listen and serve: %+v", constants.LoggerCategoryServer, err)
		}
	}()

	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.HttpServer.Shutdown(ctx)
}

func setupRouter() *gin.Engine {
	var mode = gin.ReleaseMode
	if config.ServerConfig.APIDebug {
		mode = gin.DebugMode
	}
	gin.SetMode(mode)

	router := gin.New()

	//router.Use(middlewares.CORSMiddleware())
	router.Use(gin.LoggerWithFormatter(logger.HTTPLogger))
	router.Use(gin.Recovery())

	return router
}
