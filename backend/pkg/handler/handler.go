package handler

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/yacheru/infinity-mc.ru/backend/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(cors.Default())

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
		}
	}

	return router
}
