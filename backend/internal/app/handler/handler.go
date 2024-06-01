package handler

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/yacheru/infinity-mc.ru/backend/internal/app/service"
	"github.com/yacheru/infinity-mc.ru/backend/internal/lib/api/middleware"
	"time"
)

const (
	local = "local"
	prod  = "prod"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	gin.SetMode(viper.GetString("mode"))
	router := gin.New()

	logrus.Infof("Working in %s", viper.GetString("mode"))

	var origin string
	switch viper.GetString("status") {
	case local:
		origin = "http://localhost:5173"
	case prod:
		origin = "https://infinity-mc.ru/"
	}
	_ = origin

	router.Use(cors.New(
		cors.Config{
			//AllowOrigins: []string{origin},
			AllowMethods:    []string{"GET", "POST"},
			AllowHeaders:    []string{"X-Forwarded-For", "Content-Length", "Content-Type", "Idempotence-Key"},
			AllowAllOrigins: true,
			MaxAge:          12 * time.Hour,
		}))
	router.Use(gin.Recovery())

	api := router.Group("/v1")
	{
		mc := api.Group("/mc")
		{
			mc.GET("/punishments", middleware.ValidatePunishmentsParams(), h.GetPunishments)
		}
		payment := api.Group("/payment")
		{
			payment.GET("/", middleware.ValidatePaymentParams(), h.CreatePayment)
			payment.POST("/accept", middleware.AllowedIps(), h.Accept)
		}
	}
	return router
}
