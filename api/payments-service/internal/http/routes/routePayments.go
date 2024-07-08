package routes

import (
	"github.com/gin-gonic/gin"
	"payments-service/internal/http/handlers"
	"payments-service/internal/http/middlewares"
	"payments-service/internal/kafka/producer"
)

type RoutePunishments struct {
	handler *handlers.PaymentsHandler
	router  *gin.RouterGroup
}

func NewPaymentsRoute(router *gin.RouterGroup, producer *producer.KafkaProducer) *RoutePunishments {
	handler := handlers.NewPaymentsHandler(producer)

	return &RoutePunishments{
		handler: handler,
		router:  router,
	}
}

func (r *RoutePunishments) Routes() {
	{
		r.router.GET("/create", middlewares.ValidatePaymentParams(), r.handler.CreatePayment)
		r.router.GET("/accept", r.handler.Accept)
	}
}
