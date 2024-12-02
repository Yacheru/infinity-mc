package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"jwt-auth-service/init/config"
	"jwt-auth-service/init/logger"
	"jwt-auth-service/pkg/constants"
)

func NewRedisClient(ctx context.Context, cfg *config.Config) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisHost,
		Password: cfg.RedisPassword,
		DB:       cfg.RedisDatabase,
	})

	if err := client.Ping(ctx).Err(); err != nil {
		logger.Error(err.Error(), constants.RedisCategory)
		return nil, err
	}

	return client, nil
}
