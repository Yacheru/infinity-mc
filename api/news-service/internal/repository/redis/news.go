package redis

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"news-service/init/logger"
	"news-service/internal/entities"
	"news-service/pkg/constants"
	"strconv"
	"time"
)

type NewsRedis struct {
	redis *redis.Client
	ttl   time.Duration
}

func NewNewsRedis(redis *redis.Client, ttl int) *NewsRedis {
	return &NewsRedis{
		redis: redis,
		ttl:   time.Duration(ttl) * time.Minute,
	}
}

func (r *NewsRedis) AddNews(ctx *gin.Context, discordId int, embed *entities.Embed) error {
	embedByte, err := json.Marshal(embed)
	if err != nil {
		logger.Error(err.Error(), constants.LoggerRedis)

		return err
	}

	if err := r.redis.Set(ctx.Request.Context(), strconv.Itoa(discordId), embedByte, r.ttl).Err(); err != nil {
		logger.Error(err.Error(), constants.LoggerRedis)

		return err
	}

	return nil
}

func (r *NewsRedis) GetAllNews(ctx *gin.Context) ([]*entities.Embed, error) {
	var embeds []*entities.Embed

	_, err := r.redis.Keys(ctx.Request.Context(), "*").Result()
	if err != nil {
		logger.Error(err.Error(), constants.LoggerRedis)
		return nil, err
	}

	//err = json.Unmarshal(bytes, embed)
	//if err != nil {
	//	logger.Error(err.Error(), constants.LoggerRedis)
	//	return nil, err
	//}

	return embeds, nil
}

func (r *NewsRedis) GetNewsById(ctx *gin.Context, discordId int) (*entities.Embed, error) {
	var embed = new(entities.Embed)

	bytes, err := r.redis.Get(ctx.Request.Context(), strconv.Itoa(discordId)).Bytes()
	if err != nil {
		logger.Error(err.Error(), constants.LoggerRedis)

		return nil, err
	}

	if err := json.Unmarshal(bytes, embed); err != nil {
		logger.Error(err.Error(), constants.LoggerRedis)
		return nil, err
	}

	return embed, nil
}
