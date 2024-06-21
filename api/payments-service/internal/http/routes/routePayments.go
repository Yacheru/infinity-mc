package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"payments-service/internal/http/middlewares"

	"payments-service/internal/http/handlers"
	"payments-service/internal/http/service"
	"payments-service/internal/repository/psql"
)

type RoutePunishments struct {
	handler *handlers.PaymentsHandler
	router  *gin.RouterGroup
	db      *sqlx.DB
}

func NewPaymentsRoute(router *gin.RouterGroup, db *sqlx.DB) *RoutePunishments {
	repo := psql.NewRepository(db)
	services := service.NewService(repo)
	handler := handlers.NewPaymentsHandler(services)

	return &RoutePunishments{
		handler: handler,
		router:  router,
		db:      db,
	}
}

func (r *RoutePunishments) Routes() {
	{
		r.router.GET("/create", middlewares.ValidatePaymentParams(), r.handler.CreatePayment)
		r.router.POST("/accept", r.handler.Accept)
	}
}
