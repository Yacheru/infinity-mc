package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"punishments-service/init/config"
	"punishments-service/internal/repository"
	"punishments-service/internal/server/http/handlers"
	"punishments-service/internal/server/http/middlewares"
	"punishments-service/internal/service"
)

type RoutePunishments struct {
	handler *handlers.PunishmentsHandler
	router  *gin.RouterGroup

	cfg *config.Config
}

func NewPunishmentsRoute(router *gin.RouterGroup, cfg *config.Config, db *sqlx.DB) *RoutePunishments {
	repo := repository.NewRepository(db)
	services := service.NewService(repo)
	handler := handlers.NewPunishmentsHandler(services)

	return &RoutePunishments{
		handler: handler,
		router:  router,
		cfg:     cfg,
	}
}

func (r *RoutePunishments) Routes() {
	{
		r.router.GET("/", middlewares.ValidateParams(), r.handler.GetPunishments)
	}

	admin := r.router.Group("admin", middlewares.Auth(r.cfg))
	{
		admin.GET("/unban")
		admin.GET("/unmute")
		admin.GET("/unwarn")
	}
}
