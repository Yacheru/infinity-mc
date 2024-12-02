package service

import (
	"context"
	"jwt-auth-service/init/config"
	"jwt-auth-service/internal/entities"
	"jwt-auth-service/internal/jwt"
	"jwt-auth-service/internal/repository"
	"jwt-auth-service/pkg/email"
	"jwt-auth-service/pkg/utils"
)

type AuthService interface {
	SetSession(ctx context.Context, ipAddr, userID, role string) (*entities.Tokens, error)
	RefreshTokens(ctx context.Context, refreshToken, newUserIp string) (*entities.UserResponse, error)
	RemoveToken(ctx context.Context, refreshToken string) error
}

type UserService interface {
	SendCode(ctx context.Context, email string) error
	Login(ctx context.Context, u *entities.UserLogin) (*entities.UserResponse, error)
	Register(ctx context.Context, user *entities.User, code int) (*entities.UserResponse, error)
	Logout(ctx context.Context, refreshToken string) error
	GetAll(ctx context.Context) (*[]entities.User, error)
	DeleteUser(ctx context.Context, id string) error
	UpdateRole(ctx context.Context, id, role string) error
}

type Service struct {
	AuthService
	UserService
}

func NewService(
	repo *repository.Repository,
	manager *jwt.Manager,
	email email.Sender,
	cfg *config.Config,
	hasher utils.Hasher) *Service {
	authService := NewAuthService(
		repo.AuthPostgresRepository,
		repo.UserPostgresRepository,
		repo.AuthRedisRepository,
		repo.UserRedisRepository,
		manager,
		email,
		hasher,
		cfg)

	return &Service{
		AuthService: authService,
		UserService: NewUserService(
			repo.UserPostgresRepository,
			repo.UserRedisRepository,
			repo.AuthMongoRepository,
			authService,
			email,
			hasher,
			cfg),
	}
}
