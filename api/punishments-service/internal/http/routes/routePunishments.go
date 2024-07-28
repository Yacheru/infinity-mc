package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"punishments-service/internal/http/handlers"
	"punishments-service/internal/http/middlewares"
	"punishments-service/internal/repository"
	"punishments-service/internal/service"
)

type RoutePunishments struct {
	handler *handlers.PunishmentsHandler
	router  *gin.RouterGroup
	db      *sqlx.DB
}

func NewPunishmentsRoute(router *gin.RouterGroup, db *sqlx.DB) *RoutePunishments {
	repo := repository.NewRepository(db)
	services := service.NewService(repo)
	handler := handlers.NewPunishmentsHandler(services)

	return &RoutePunishments{
		handler: handler,
		router:  router,
		db:      db,
	}
}

func (r *RoutePunishments) Routes() {
	{
		r.router.GET("/", middlewares.ValidateParams(), r.handler.GetPunishments)
	}
}
