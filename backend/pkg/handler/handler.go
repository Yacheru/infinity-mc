package handler

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/yacheru/infinity-mc.ru/backend/pkg/service"
	"time"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(cors.New(
		cors.Config{
			AllowMethods:    []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
			AllowHeaders:    []string{"Origin", "Content-Length", "Content-Type", "Authorization", "Idempotence-Key"},
			AllowAllOrigins: true,
			MaxAge:          12 * time.Hour,
		}))

	api := router.Group("/v1", h.userIdentity)
	{
		mc := api.Group("/mc")
		{
			mc.GET("/", h.Mc)
			mc.GET("/bans", h.GetAllBans)
		}
		payment := api.Group("/payment")
		{
			payment.GET("/", h.CreatePayment)
		}
	}
	return router
}
