package service

import (
	"context"
	"database/sql"
	"errors"
	"jwt-auth-service/internal/jwt"
	hash "jwt-auth-service/pkg/utils"
	"time"

	"jwt-auth-service/init/config"
	"jwt-auth-service/init/logger"
	"jwt-auth-service/internal/entities"
	"jwt-auth-service/internal/repository"
	"jwt-auth-service/pkg/constants"
	"jwt-auth-service/pkg/email"
)

type Auth struct {
	authPostgres repository.AuthPostgresRepository
	userPostgres repository.UserPostgresRepository
	authRedis    repository.AuthRedisRepository
	userRedis    repository.UserRedisRepository

	email    email.Sender
	hasher   hash.Hasher
	tManager jwt.TokenManager

	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration
}

func NewAuthService(
	authPostgres repository.AuthPostgresRepository,
	userPostgres repository.UserPostgresRepository,
	authRedis repository.AuthRedisRepository,
	userRedis repository.UserRedisRepository,
	tManager jwt.TokenManager,
	email email.Sender,
	hasher hash.Hasher,
	cfg *config.Config) *Auth {
	return &Auth{
		authPostgres:    authPostgres,
		userPostgres:    userPostgres,
		authRedis:       authRedis,
		userRedis:       userRedis,
		email:           email,
		hasher:          hasher,
		tManager:        tManager,
		accessTokenTTL:  time.Duration(cfg.AccessTokenTTL) * time.Minute,
		refreshTokenTTL: time.Duration(cfg.RefreshTokenTTL) * time.Minute,
	}
}

func (a *Auth) SetSession(ctx context.Context, ipAddr, userID, role string) (*entities.Tokens, error) {
	logger.Debug("setting up a new session", ipAddr)

	var (
		tokens = new(entities.Tokens)
		err    error
	)

	tokens.AccessToken, err = a.tManager.NewAccessToken(ipAddr, userID, role, a.accessTokenTTL)
	if err != nil {
		logger.Error(err.Error(), constants.ServiceCategory)
		return nil, err
	}

	logger.Debug("access token generated", ipAddr)

	tokens.RefreshToken, err = a.tManager.NewRefreshToken(ipAddr, userID, role, a.refreshTokenTTL)
	if err != nil {
		logger.Error(err.Error(), constants.ServiceCategory)
		return nil, err
	}

	logger.Debug("refresh token generated", ipAddr)

	session := &entities.Session{
		RefreshToken: tokens.RefreshToken,
		ExpiresIn:    time.Now().Add(a.refreshTokenTTL).Unix(),
	}

	if err := a.authPostgres.SetToken(ctx, userID, session); err != nil {
		return nil, err
	}

	logger.Debug("success postgres SetSession", ipAddr)

	if err := a.authRedis.SetSession(ctx, userID, session); err != nil {
		return nil, err
	}

	logger.Debug("success redis SetSession", ipAddr)

	return tokens, nil
}

func (a *Auth) RefreshTokens(ctx context.Context, refreshToken, newUserIp string) (*entities.UserResponse, error) {
	logger.DebugF("validate refreshToken (%s)", newUserIp, refreshToken)

	if _, err := a.tManager.ValidToken(refreshToken); err != nil {
		if !errors.Is(err, jwt.TokenExpired) {
			logger.Error(err.Error(), constants.ServiceCategory)
			return nil, constants.UnauthorizedError
		}
		logger.Error(err.Error(), constants.ServiceCategory)
	}

	logger.Debug("refreshToken valid", newUserIp)

	user, err := a.userPostgres.GetUserByRefresh(ctx, refreshToken)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, constants.UserNotFoundError
		}
		return nil, err
	}
	if user == nil {
		return nil, constants.UnauthorizedError
	}

	logger.DebugF("successful receipt user by email (%+v)", newUserIp, user)

	tokens, err := a.SetSession(ctx, newUserIp, user.UserID, user.Role)
	if err != nil {
		return nil, err
	}

	logger.DebugF("successful receipt of tokens (%+v)", newUserIp, tokens)

	userResponse := &entities.UserResponse{
		User:   *user,
		Tokens: *tokens,
	}

	return userResponse, nil
}

func (a *Auth) RemoveToken(ctx context.Context, refreshToken string) error {
	if err := a.authPostgres.RemoveToken(ctx, refreshToken); err != nil {
		return err
	}
	return nil
}
