package service

import (
	"context"
	"websocket-chat-service/internal/entities"
	"websocket-chat-service/internal/repository"
)

type ScyllaService interface {
	StoreMessage(ctx context.Context, msg *entities.Message) error
	GetAllMessages(ctx context.Context) ([]*entities.Message, error)
	GetPlayerMessages(ctx context.Context, nickname string, limit, offset int) ([]*entities.Message, error)
}

type Service struct {
	ScyllaService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		ScyllaService: NewMessagesService(repo.ScyllaRepository, repo.RedisRepository),
	}
}
