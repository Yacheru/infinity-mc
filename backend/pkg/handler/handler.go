package handler

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Handler struct {
}

//func NewHandler(services *service.Service) *Handler {
//	return &Handler{services: services}
//}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(cors.Default())

	api := router.Group("/v1", h.ProtectApi)
	{
		mc := api.Group("/mc")
		{
			mc.GET("/", h.Mc)
			mc.POST("/payments", h.CreatePayment)
		}
	}

	return router
}
