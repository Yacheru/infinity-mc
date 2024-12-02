package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"websocket-chat-service/internal/entities"
)

type MessageRepository struct {
	client *redis.Client
}

func NewMessageRepository(client *redis.Client) *MessageRepository {
	return &MessageRepository{client}
}

func (m *MessageRepository) StoreMessage(ctx context.Context, msg *entities.Message) error {
	return nil
}
