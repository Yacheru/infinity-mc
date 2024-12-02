package service

import (
	"context"

	"websocket-chat-service/internal/entities"
	"websocket-chat-service/internal/repository"
)

type MessagesService struct {
	scylla repository.ScyllaRepository
	redis  repository.RedisRepository
}

func NewMessagesService(scylla repository.ScyllaRepository, redis repository.RedisRepository) *MessagesService {
	return &MessagesService{scylla: scylla, redis: redis}
}

func (s *MessagesService) GetPlayerMessages(ctx context.Context, nickname string, limit, offset int) ([]*entities.Message, error) {
	messages, err := s.scylla.GetPlayerMessages(ctx, nickname, limit, offset)
	if err != nil {
		return nil, err
	}

	return messages, nil
}

func (s *MessagesService) GetAllMessages(ctx context.Context) ([]*entities.Message, error) {
	messages, err := s.scylla.GetAllMessages(ctx)
	if err != nil {
		return nil, err
	}

	return messages, nil
}

func (s *MessagesService) StoreMessage(ctx context.Context, msg *entities.Message) error {
	if err := s.scylla.StoreMessage(ctx, msg); err != nil {
		return err
	}

	if err := s.redis.StoreMessage(ctx, msg); err != nil {
		return err
	}

	return nil
}
