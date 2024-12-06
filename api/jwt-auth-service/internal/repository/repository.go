package repository

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"

	"jwt-auth-service/init/config"
	"jwt-auth-service/internal/entities"
	"jwt-auth-service/internal/repository/mongodb"
	"jwt-auth-service/internal/repository/postgres"
	r "jwt-auth-service/internal/repository/redis"
)

//go:generate go run github.com/vektra/mockery/v2@v2.49.2 --name=AuthPostgresRepository
type AuthPostgresRepository interface {
	SetToken(ctx context.Context, userId string, session *entities.Session) error
	RemoveToken(ctx context.Context, refreshToken string) error
}

//go:generate go run github.com/vektra/mockery/v2@v2.49.2 --name=UserPostgresRepository
type UserPostgresRepository interface {
	StoreNewUser(ctx context.Context, u *entities.User) (*entities.User, error)
	GetUserByRefresh(ctx context.Context, refreshToken string) (*entities.User, error)
	GetUserID(ctx context.Context, email, password string) (string, error)
	GetUserByEmail(ctx context.Context, email string) (*entities.User, error)
	GetAll(ctx context.Context) (*[]entities.User, error)
	DeleteUser(ctx context.Context, id string) error
	UpdateRole(ctx context.Context, id, role string) error
}

//go:generate go run github.com/vektra/mockery/v2@v2.49.2 --name=AuthRedisRepository
type AuthRedisRepository interface {
	SetSession(ctx context.Context, userId string, session *entities.Session) error
}

//go:generate go run github.com/vektra/mockery/v2@v2.49.2 --name=UserRedisRepository
type UserRedisRepository interface {
	StoreNewUser(ctx context.Context, u *entities.User) error
	GetUserById(ctx context.Context, userId string) (*entities.User, error)
}

//go:generate go run github.com/vektra/mockery/v2@v2.49.2 --name=AuthMongoRepository
type AuthMongoRepository interface {
	SetCode(ctx context.Context, email string, code int, expiration time.Time) error
	GetCode(ctx context.Context, email string, code int) error
}

type Repository struct {
	AuthPostgresRepository
	UserPostgresRepository
	AuthRedisRepository
	UserRedisRepository
	AuthMongoRepository
}

func NewRepository(pdb *sqlx.DB, redis *redis.Client, coll *mongo.Collection, cfg *config.Config) *Repository {
	return &Repository{
		AuthPostgresRepository: postgres.NewAuthPostgres(pdb),
		UserPostgresRepository: postgres.NewUserPostgres(pdb),
		AuthRedisRepository:    r.NewAuthRedis(redis),
		UserRedisRepository:    r.NewUserRedis(redis, cfg.RedisTTL),
		AuthMongoRepository:    mongodb.NewAuthMongo(coll),
	}
}
