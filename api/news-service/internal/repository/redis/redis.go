package redis

import (
	"context"
	"news-service/init/logger"
	"news-service/pkg/constants"

	"github.com/redis/go-redis/v9"

	"news-service/init/config"
)

func NewRedisClient(ctx context.Context, cfg *config.Config) (*redis.Client, error) {
	logger.Debug("creating a new Redis client...", constants.LoggerRedis)

	opt := &redis.Options{
		Addr:     cfg.RedisHost,
		Password: cfg.RedisPassword,
		DB:       0,
	}

	client := redis.NewClient(opt)

	logger.Debug("new Redis client successfully created. Pinging...", constants.LoggerRedis)

	_, err := client.Ping(ctx).Result()
	if err != nil {
		logger.Error(err.Error(), constants.LoggerRedis)

		return nil, err
	}

	logger.Info("redis client is working", constants.LoggerRedis)

	return client, nil
}
