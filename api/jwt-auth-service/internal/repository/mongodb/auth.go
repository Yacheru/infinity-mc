package mongodb

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"jwt-auth-service/init/logger"
	"jwt-auth-service/pkg/constants"
	"time"

	"go.mongodb.org/mongo-driver/mongo"

	"jwt-auth-service/internal/entities"
)

type AuthMongo struct {
	coll *mongo.Collection
}

func NewAuthMongo(coll *mongo.Collection) *AuthMongo {
	return &AuthMongo{coll: coll}
}

func (m *AuthMongo) GetCode(ctx context.Context, email string, code int) error {
	if err := m.coll.FindOne(ctx, bson.M{"code": code, "email": email}).Err(); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return constants.CodeNotFoundError
		}

		logger.Error(err.Error(), constants.MongoCategory)
		return err
	}

	return nil
}

func (m *AuthMongo) SetCode(ctx context.Context, email string, code int, expiration time.Time) error {
	emailVerifEntity := &entities.EmailVerification{
		Email:      email,
		Code:       code,
		Expiration: expiration,
	}

	_, bytes, err := bson.MarshalValue(emailVerifEntity)
	if err != nil {
		logger.Error(err.Error(), constants.MongoCategory)
		return err
	}

	if _, err := m.coll.InsertOne(ctx, bytes); err != nil {
		logger.Error(err.Error(), constants.MongoCategory)
		return err
	}

	return nil
}
