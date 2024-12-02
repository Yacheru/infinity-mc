package repository

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/scylladb/gocqlx/v2"
	"websocket-chat-service/internal/entities"
	r "websocket-chat-service/internal/repository/redis"
	"websocket-chat-service/internal/repository/scylladb"
)

type ScyllaRepository interface {
	StoreMessage(ctx context.Context, msg *entities.Message) error
	GetAllMessages(ctx context.Context) ([]*entities.Message, error)
	GetPlayerMessages(ctx context.Context, nickname string, limit, offset int) ([]*entities.Message, error)
}

type RedisRepository interface {
	StoreMessage(ctx context.Context, msg *entities.Message) error
}

type Repository struct {
	ScyllaRepository
	RedisRepository
}

func NewRepository(scylla *gocqlx.Session, redis *redis.Client) *Repository {
	return &Repository{
		ScyllaRepository: scylladb.NewMessageRepository(scylla),
		RedisRepository:  r.NewMessageRepository(redis),
	}
}
