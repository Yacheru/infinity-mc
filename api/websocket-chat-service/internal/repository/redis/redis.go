package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"websocket-chat-service/init/config"
	"websocket-chat-service/init/logger"
	"websocket-chat-service/pkg/constants"
)

func NewRedisClient(ctx context.Context, cfg *config.Config) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: cfg.RedisPassword,
		DB:       0,
	})

	if err := client.Ping(ctx).Err(); err != nil {
		logger.Error(err.Error(), "Ping: "+constants.RedisCategory)
		return nil, err
	}

	return client, nil
}
