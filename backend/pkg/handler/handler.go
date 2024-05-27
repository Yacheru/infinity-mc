package handler

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
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
	gin.SetMode(viper.GetString("mode"))
	router := gin.New()

	logrus.Infof("Working in %s", viper.GetString("mode"))

	var origin string
	_ = origin
	switch viper.GetString("status") {
	case "local":
		origin = "http://localhost:5173"
	case "prod":
		origin = "https://infinity-mc.ru/"
	}

	router.Use(cors.New(
		cors.Config{
			//AllowOrigins: []string{origin},
			AllowMethods:    []string{"GET", "POST"},
			AllowHeaders:    []string{"X-Forwarded-For", "Content-Length", "Content-Type", "Authorization", "Idempotence-Key"},
			AllowAllOrigins: true,
			MaxAge:          12 * time.Hour,
		}))

	api := router.Group("/v1")
	{
		mc := api.Group("/mc")
		{
			mc.GET("/", h.Mc)
			mc.GET("/bans", h.GetAllBans)
		}
		payment := api.Group("/payment")
		{
			payment.GET("/", h.CreatePayment)
			payment.POST("/accept", h.Accept)
		}
	}
	return router
}
