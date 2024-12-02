package redis

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"jwt-auth-service/init/logger"
	"jwt-auth-service/internal/entities"
	"jwt-auth-service/pkg/constants"
	"time"
)

type UserRedis struct {
	client *redis.Client
	ttl    time.Duration
}

func NewUserRedis(client *redis.Client, ttl int) *UserRedis {
	return &UserRedis{
		client: client,
		ttl:    time.Duration(ttl) * time.Minute,
	}
}

func (u *UserRedis) StoreNewUser(ctx context.Context, user *entities.User) error {
	byteUser, err := json.Marshal(user)
	if err != nil {
		logger.Error(err.Error(), constants.RedisCategory)

		return err
	}

	if err := u.client.Set(ctx, user.UserID, byteUser, u.ttl).Err(); err != nil {
		logger.Error(err.Error(), constants.RedisCategory)

		return err
	}

	return nil
}

func (u *UserRedis) GetUserById(ctx context.Context, userId string) (*entities.User, error) {
	var user = new(entities.User)

	bytes, err := u.client.Get(ctx, userId).Bytes()
	if err != nil {
		logger.Error(err.Error(), constants.RedisCategory)

		return nil, err
	}

	if err := json.Unmarshal(bytes, user); err != nil {
		logger.Error(err.Error(), constants.RedisCategory)

		return nil, err
	}

	return user, nil
}
