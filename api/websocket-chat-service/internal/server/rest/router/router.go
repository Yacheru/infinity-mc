package router

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/scylladb/gocqlx/v2"
	"websocket-chat-service/init/config"
	"websocket-chat-service/internal/repository"
	"websocket-chat-service/internal/server/rest/handlers"
	"websocket-chat-service/internal/server/rest/middlewares"
	"websocket-chat-service/internal/server/websocket"
	"websocket-chat-service/internal/service"
)

type Router struct {
	router  *gin.RouterGroup
	handler *handlers.Handler
}

func NewRouterAndComponents(cfg *config.Config, router *gin.RouterGroup, scylla *gocqlx.Session, r *redis.Client) *Router {
	repo := repository.NewRepository(scylla, r)
	services := service.NewService(repo)
	ws := websocket.NewWebSocket(cfg, services)
	handler := handlers.NewHandler(services, ws)

	return &Router{
		router:  router,
		handler: handler,
	}
}

func (r *Router) Router() {
	ws := r.router.Group("/ws")
	{
		ws.GET("/run", r.handler.RunWS)
		ws.GET("/close", r.handler.CloseWS)
	}
	messages := r.router.Group("/messages")
	{
		messages.GET("/all", r.handler.GetAllMessages)
		messages.GET("/player/:nickname", middlewares.ValidatePagination(), r.handler.GetPlayerMessages)
	}
}
