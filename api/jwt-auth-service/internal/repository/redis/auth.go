package redis

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"jwt-auth-service/init/logger"
	"jwt-auth-service/internal/entities"
	"jwt-auth-service/pkg/constants"
)

type AuthRedis struct {
	client *redis.Client
}

func NewAuthRedis(client *redis.Client) *AuthRedis {
	return &AuthRedis{client: client}
}

func (r *AuthRedis) SetSession(ctx context.Context, userId string, session *entities.Session) error {
	var user = new(entities.User)

	bytes, err := r.client.Get(ctx, userId).Bytes()
	if err != nil {
		logger.Error(err.Error(), constants.RedisCategory)
		return err
	}

	err = json.Unmarshal(bytes, user)
	if err != nil {
		logger.Error(err.Error(), constants.RedisCategory)
		return err
	}

	userBytes, err := json.Marshal(user)
	if err != nil {
		logger.Error(err.Error(), constants.RedisCategory)
		return err
	}

	if err := r.client.Set(ctx, userId, userBytes, -1).Err(); err != nil {
		logger.Error(err.Error(), constants.RedisCategory)

		return err
	}

	return nil
}
