package service

import (
	"context"
	"errors"
	"fmt"
	"jwt-auth-service/init/config"
	"jwt-auth-service/init/logger"
	"jwt-auth-service/internal/utils"
	"time"

	"jwt-auth-service/internal/entities"
	"jwt-auth-service/internal/repository"
	"jwt-auth-service/pkg/constants"
	"jwt-auth-service/pkg/email"
	pkgUtils "jwt-auth-service/pkg/utils"
)

type User struct {
	userPostgres repository.UserPostgresRepository
	userRedis    repository.UserRedisRepository
	authMongo    repository.AuthMongoRepository

	authService AuthService

	email  email.Sender
	hasher pkgUtils.Hasher

	cfg *config.Config
}

func NewUserService(
	userPostgres repository.UserPostgresRepository,
	userRedis repository.UserRedisRepository,
	authMongo repository.AuthMongoRepository,
	authService AuthService,
	email email.Sender,
	hasher pkgUtils.Hasher,
	cfg *config.Config) *User {
	return &User{
		userPostgres: userPostgres,
		userRedis:    userRedis,
		authMongo:    authMongo,
		authService:  authService,
		email:        email,
		hasher:       hasher,
		cfg:          cfg,
	}
}

func (a *User) SendCode(ctx context.Context, email string) error {
	expiration := time.Now().Add(time.Duration(a.cfg.MongoTTL) * time.Second)
	code := utils.GenerateVerifCode()

	user, err := a.userPostgres.GetUserByEmail(ctx, email)
	if err != nil {
		if !errors.Is(err, constants.UserNotFoundError) {
			return err
		}
	}
	if user != nil {
		return constants.UserAlreadyExistsError
	}

	if err := a.authMongo.SetCode(ctx, email, code, expiration); err != nil {
		return err
	}

	message := fmt.Sprintf("<b>%d</b> - your verification code", code)

	if err := a.email.SendMail([]string{email}, []byte(message), "Verification code"); err != nil {
		logger.Error(err.Error(), constants.ServiceCategory)
		return err
	}

	return nil
}

func (a *User) Register(ctx context.Context, user *entities.User, code int) (*entities.UserResponse, error) {
	if err := a.authMongo.GetCode(ctx, user.Email, code); err != nil {
		return nil, err
	}

	user.Password = a.hasher.Hash(user.Password)

	userdb, err := a.userPostgres.StoreNewUser(ctx, user)
	if err != nil {
		return nil, err
	}
	if err := a.userRedis.StoreNewUser(ctx, user); err != nil {
		return nil, err
	}

	tokens, err := a.authService.SetSession(ctx, user.IpAddr, userdb.UserID, userdb.Role)
	if err != nil {
		return nil, err
	}

	if err := a.email.SendMail([]string{user.Email}, []byte("You have successfully registered"), "Registry in infinity-mc"); err != nil {
		logger.Error(err.Error(), constants.ServiceCategory)
	}

	userResponse := &entities.UserResponse{
		User:   *userdb,
		Tokens: *tokens,
	}

	return userResponse, nil
}

func (a *User) Login(ctx context.Context, u *entities.UserLogin) (*entities.UserResponse, error) {
	user, err := a.userPostgres.GetUserByEmail(ctx, u.Email)
	if err != nil {
		if errors.Is(err, constants.UserNotFoundError) {
			return nil, constants.UserNotFoundError
		}
		return nil, err
	}
	u.Password = a.hasher.Hash(u.Password)

	if user.Password != u.Password {
		return nil, constants.UserNotFoundError
	}

	tokens, err := a.authService.SetSession(ctx, u.IpAddress, user.UserID, user.Role)
	if err != nil {
		return nil, err
	}

	userResponse := &entities.UserResponse{
		User:   *user,
		Tokens: *tokens,
	}

	return userResponse, nil
}

func (a *User) Logout(ctx context.Context, refreshToken string) error {
	if err := a.authService.RemoveToken(ctx, refreshToken); err != nil {
		return err
	}
	return nil
}

func (a *User) GetAll(ctx context.Context) (*[]entities.User, error) {
	users, err := a.userPostgres.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (a *User) DeleteUser(ctx context.Context, id string) error {
	if err := a.userPostgres.DeleteUser(ctx, id); err != nil {
		return err
	}
	return nil
}

func (a *User) UpdateRole(ctx context.Context, id, role string) error {
	if err := a.userPostgres.UpdateRole(ctx, id, role); err != nil {
		return err
	}
	return nil
}
