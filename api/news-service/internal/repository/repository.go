package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
	"news-service/init/config"

	"news-service/internal/entities"
	"news-service/internal/repository/elastic"
	"news-service/internal/repository/postgres"
	r "news-service/internal/repository/redis"
)

type NewsPostgres interface {
	AddNew(ctx *gin.Context, discordId int, embed *entities.Embed) (*int, error)
	GetAllNews(ctx *gin.Context) ([]*entities.Embed, error)
	GetNewsById(ctx *gin.Context, discordId int) (*entities.Embed, error)
}

type NewsElastic interface {
}

type NewsRedis interface {
	AddNews(ctx *gin.Context, discordId int, embed *entities.Embed) error
	GetAllNews(ctx *gin.Context) ([]*entities.Embed, error)
	GetNewsById(ctx *gin.Context, discordId int) (*entities.Embed, error)
}

type Repository struct {
	NewsPostgres
	NewsElastic
	NewsRedis
}

func NewRepository(db *sqlx.DB, es *elastic.Client, redis *redis.Client, cfg *config.Config) *Repository {
	return &Repository{
		NewsPostgres: postgres.NewNewsPostgres(db),
		NewsElastic:  elastic.NewNewsElastic(es),
		NewsRedis:    r.NewNewsRedis(redis, cfg.RedisTTL),
	}
}
