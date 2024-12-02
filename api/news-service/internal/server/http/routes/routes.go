package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"

	"news-service/init/config"
	"news-service/internal/discord"
	"news-service/internal/repository"
	"news-service/internal/repository/elastic"
	"news-service/internal/server/http/handlers"
	"news-service/internal/server/http/middlewares"
	"news-service/internal/service"
)

type Router struct {
	router  *gin.RouterGroup
	handler *handlers.Handlers
}

func NewComponentsAndRoutes(router *gin.RouterGroup, db *sqlx.DB, es *elastic.Client, redis *redis.Client, cfg *config.Config) *Router {
	repo := repository.NewRepository(db, es, redis, cfg)
	webhook := discord.NewWebhookClient(cfg)
	services := service.NewService(repo, webhook)
	handler := handlers.NewHandlers(services)

	return &Router{router: router, handler: handler}
}

func (r *Router) Routes() {
	{
		r.router.POST("/add", r.handler.AddNew)
		r.router.GET("/all", r.handler.GetAllNews)
		r.router.GET("/:id", middlewares.ParseParam(), r.handler.GetNewsById)
	}
}
