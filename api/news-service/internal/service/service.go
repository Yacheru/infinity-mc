package service

import (
	"github.com/gin-gonic/gin"
	"news-service/internal/discord"

	"news-service/internal/entities"
	"news-service/internal/repository"
	"news-service/internal/service/news"
)

type News interface {
	AddNew(ctx *gin.Context, embed *entities.Embed) (*int, error)
	GetAllNews(ctx *gin.Context) ([]*entities.Embed, error)
	GetNewsById(ctx *gin.Context, discordId int) (*entities.Embed, error)
}

type Service struct {
	News
}

func NewService(repo *repository.Repository, ds *discord.WebhookClient) *Service {
	return &Service{
		News: news.NewNewsService(repo.NewsPostgres, repo.NewsRedis, ds),
	}
}
