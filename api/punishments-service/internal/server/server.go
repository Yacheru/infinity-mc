package server

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"punishments-service/internal/repository/mysql"
	"punishments-service/internal/server/http/middlewares"
	"punishments-service/internal/server/http/routes"
	"punishments-service/pkg/constants"
	"time"

	"github.com/gin-gonic/gin"
	"punishments-service/init/config"
	"punishments-service/init/logger"
)

type Server struct {
	HttpServer *http.Server
}

func NewServer(ctx context.Context, cfg *config.Config) (*Server, error) {
	db, err := mysql.InitMySQLDatabase(ctx, cfg)
	if err != nil {
		return nil, err
	}

	router := setupRouter()

	api := router.Group(cfg.ApiEntry)
	routes.NewPunishmentsRoute(api, cfg, db).Routes()

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

func (s *Server) Run(ctx context.Context) error {
	errs, ctx := errgroup.WithContext(ctx)

	errs.Go(func() error {
		if err := s.HttpServer.ListenAndServe(); err != nil && !errors.Is(http.ErrServerClosed, err) {
			return err
		}
		return nil
	})

	if err := errs.Wait(); err != nil {
		logger.ErrorF("failed to listen and serve: %s", constants.LoggerCategoryServer, err.Error())
		return err
	}
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

	router.Use(gin.Recovery())
	router.Use(gin.LoggerWithFormatter(logger.HTTPLogger))
	router.Use(middlewares.RateLimiter())
	router.Use(middlewares.CORSMiddleware())

	return router
}
