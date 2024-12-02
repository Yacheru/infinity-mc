package server

import (
	"context"
	"fmt"
	"jwt-auth-service/internal/repository/redis"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"jwt-auth-service/init/config"
	"jwt-auth-service/init/logger"
	"jwt-auth-service/internal/repository/mongodb"
	"jwt-auth-service/internal/repository/postgres"
	"jwt-auth-service/internal/server/http/middleware"
	"jwt-auth-service/internal/server/http/routes"
)

type HTTPServer struct {
	server *http.Server
}

func NewServer(ctx context.Context, cfg *config.Config, log *logrus.Logger) (*HTTPServer, error) {
	pdb, err := postgres.NewPostgresConnection(ctx, cfg, log)
	if err != nil {
		return nil, err
	}

	coll, err := mongodb.InitMongoDB(ctx, cfg)
	if err != nil {
		return nil, err
	}

	r, err := redis.NewRedisClient(ctx, cfg)
	if err != nil {
		return nil, err
	}

	engine := setupGin(cfg)
	router := engine.Group(cfg.ApiEntry)
	routes.InitRouterAndComponents(router, pdb, r, coll, cfg).Routes()

	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", cfg.ApiPort),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		Handler:        engine,
		MaxHeaderBytes: 1 << 20,
	}

	return &HTTPServer{server: server}, nil
}

func (s *HTTPServer) Run() error {
	return s.server.ListenAndServe()
}

func (s *HTTPServer) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

func setupGin(cfg *config.Config) *gin.Engine {
	var mode = gin.ReleaseMode
	if cfg.ApiDebug {
		mode = gin.DebugMode
	}
	gin.SetMode(mode)

	engine := gin.New()

	engine.Use(middleware.CORSMiddleware())
	engine.Use(gin.Recovery())
	engine.Use(gin.LoggerWithFormatter(logger.HTTPLogger))

	return engine
}
