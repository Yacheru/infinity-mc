package server

import (
	"context"
	"errors"
	"fmt"
	"go.elastic.co/apm/module/apmgin/v2"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"news-service/init/config"
	"news-service/init/logger"
	"news-service/internal/repository/elastic"
	"news-service/internal/repository/postgres"
	"news-service/internal/repository/redis"
	"news-service/internal/server/http/routes"
	"news-service/pkg/constants"
)

type HTTPServer struct {
	server *http.Server
}

func NewServer(ctx context.Context, cfg *config.Config, log *logrus.Logger) (*HTTPServer, error) {
	db, err := postgres.NewPostgresConnection(ctx, cfg, log)
	if err != nil {
		return nil, err
	}

	es, err := elastic.NewElasticClient(cfg)
	if err != nil {
		return nil, err
	}

	r, err := redis.NewRedisClient(ctx, cfg)
	if err != nil {
		return nil, err
	}

	engine := SetupGin(cfg)
	api := engine.Group(cfg.APIEntry)
	routes.NewComponentsAndRoutes(api, db, es, r, cfg).Routes()

	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", cfg.APIPort),
		Handler:        engine,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return &HTTPServer{
		server: server,
	}, nil
}

func (s *HTTPServer) Run(cfg *config.Config) error {
	go func() {
		if err := s.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error(err.Error(), constants.LoggerServer)
		}
	}()

	logger.InfoF("success to listen and serve on :%d port", constants.LoggerServer, cfg.APIPort)

	return nil
}

func (s *HTTPServer) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

func SetupGin(cfg *config.Config) *gin.Engine {
	var mode = gin.ReleaseMode
	if cfg.APIDebug {
		mode = gin.DebugMode
	}
	gin.SetMode(mode)

	engine := gin.New()

	engine.Use(gin.Recovery())
	engine.Use(gin.LoggerWithFormatter(logger.HTTPLogger))
	engine.Use(apmgin.Middleware(engine))

	return engine
}
